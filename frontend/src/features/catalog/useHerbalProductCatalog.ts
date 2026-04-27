import { useCallback, useEffect, useMemo, useState } from 'react'
import type {
  CatalogFetchState,
  CatalogSortPreference,
  HerbalCatalogProduct,
} from './catalogTypes'

const CATALOG_PRODUCTS_ENDPOINT =
  import.meta.env.VITE_CATALOG_PRODUCTS_ENDPOINT ?? '/api/catalog'

const friendlyCatalogLoadError =
  'Katalog herbal belum bisa dimuat. Periksa koneksi atau coba lagi sebentar.'

function isHerbalCatalogProduct(candidate: unknown): candidate is HerbalCatalogProduct {
  if (!candidate || typeof candidate !== 'object') {
    return false
  }

  const product = candidate as Partial<Record<keyof HerbalCatalogProduct, unknown>>

  return (
    typeof product.id === 'string' &&
    typeof product.name === 'string' &&
    typeof product.slug === 'string' &&
    typeof product.price === 'number' &&
    typeof product.stock_quantity === 'number' &&
    typeof product.image_url === 'string' &&
    typeof product.is_active === 'boolean'
  )
}

function extractCatalogProducts(apiPayload: unknown): HerbalCatalogProduct[] {
  if (Array.isArray(apiPayload)) {
    return apiPayload.filter(isHerbalCatalogProduct)
  }

  if (!apiPayload || typeof apiPayload !== 'object') {
    return []
  }

  const possiblePayload = apiPayload as {
    data?: unknown
    products?: unknown
    items?: unknown
  }

  const productCollection =
    possiblePayload.data ?? possiblePayload.products ?? possiblePayload.items

  return Array.isArray(productCollection)
    ? productCollection.filter(isHerbalCatalogProduct)
    : []
}

function matchesCatalogSearch(product: HerbalCatalogProduct, normalizedSearchTerm: string) {
  if (!normalizedSearchTerm) {
    return true
  }

  const searchableProductCopy = [
    product.name,
    product.description,
    product.benefit,
    product.composition,
    product.manufacturer,
    product.regency,
  ]
    .join(' ')
    .toLowerCase()

  return searchableProductCopy.includes(normalizedSearchTerm)
}

function sortCatalogProducts(
  products: HerbalCatalogProduct[],
  sortPreference: CatalogSortPreference,
) {
  return [...products].sort((leftProduct, rightProduct) => {
    if (sortPreference === 'price-ascending') {
      return leftProduct.price - rightProduct.price
    }

    if (sortPreference === 'rating-descending') {
      return rightProduct.average_rating - leftProduct.average_rating
    }

    return Number(rightProduct.is_active) - Number(leftProduct.is_active)
  })
}

export function useHerbalProductCatalog() {
  const [catalogFetchState, setCatalogFetchState] = useState<CatalogFetchState>({
    products: [],
    isLoading: true,
    errorMessage: null,
  })
  const [catalogSearchTerm, setCatalogSearchTerm] = useState('')
  const [selectedRegency, setSelectedRegency] = useState('all')
  const [sortPreference, setSortPreference] =
    useState<CatalogSortPreference>('featured')

  const loadCatalogProducts = useCallback(async (signal?: AbortSignal) => {
    await Promise.resolve()

    setCatalogFetchState((currentState) => ({
      ...currentState,
      isLoading: true,
      errorMessage: null,
    }))

    try {
      const catalogResponse = await fetch(CATALOG_PRODUCTS_ENDPOINT, {
        headers: { Accept: 'application/json' },
        signal,
      })

      if (!catalogResponse.ok) {
        throw new Error(`Catalog API responded with ${catalogResponse.status}`)
      }

      const catalogPayload: unknown = await catalogResponse.json()
      const activeCatalogProducts = extractCatalogProducts(catalogPayload).filter(
        (product) => product.is_active,
      )

      setCatalogFetchState({
        products: activeCatalogProducts,
        isLoading: false,
        errorMessage: null,
      })
    } catch (error) {
      if (error instanceof DOMException && error.name === 'AbortError') {
        return
      }

      setCatalogFetchState({
        products: [],
        isLoading: false,
        errorMessage: friendlyCatalogLoadError,
      })
    }
  }, [])

  useEffect(() => {
    const catalogRequestController = new AbortController()
    const initialCatalogRequestId = window.setTimeout(() => {
      void loadCatalogProducts(catalogRequestController.signal)
    }, 0)

    return () => {
      window.clearTimeout(initialCatalogRequestId)
      catalogRequestController.abort()
    }
  }, [loadCatalogProducts])

  const availableRegencies = useMemo(() => {
    const uniqueRegencies = new Set(
      catalogFetchState.products
        .map((product) => product.regency)
        .filter((regency) => regency.trim().length > 0),
    )

    return Array.from(uniqueRegencies).sort((leftRegency, rightRegency) =>
      leftRegency.localeCompare(rightRegency),
    )
  }, [catalogFetchState.products])

  const visibleCatalogProducts = useMemo(() => {
    const normalizedSearchTerm = catalogSearchTerm.trim().toLowerCase()
    const filteredProducts = catalogFetchState.products.filter((product) => {
      const matchesRegency =
        selectedRegency === 'all' || product.regency === selectedRegency

      return matchesRegency && matchesCatalogSearch(product, normalizedSearchTerm)
    })

    return sortCatalogProducts(filteredProducts, sortPreference)
  }, [catalogFetchState.products, catalogSearchTerm, selectedRegency, sortPreference])

  return {
    availableRegencies,
    catalogSearchTerm,
    errorMessage: catalogFetchState.errorMessage,
    isLoading: catalogFetchState.isLoading,
    productsTotal: catalogFetchState.products.length,
    retryCatalogRequest: () => void loadCatalogProducts(),
    selectedRegency,
    setCatalogSearchTerm,
    setSelectedRegency,
    setSortPreference,
    sortPreference,
    visibleCatalogProducts,
  }
}
