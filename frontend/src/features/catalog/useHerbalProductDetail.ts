import { useCallback, useEffect, useMemo, useState } from 'react'
import type { HerbalCatalogProduct } from './catalogTypes'
import { extractCatalogProductDetail } from './catalogProductGuards'

const CATALOG_PRODUCT_DETAIL_ENDPOINT =
  import.meta.env.VITE_CATALOG_PRODUCT_DETAIL_ENDPOINT ?? '/api/catalog/:slug'

const friendlyProductDetailError =
  'Detail produk belum bisa dimuat. Periksa koneksi atau coba lagi sebentar.'

interface ProductDetailFetchState {
  product: HerbalCatalogProduct | null
  isLoading: boolean
  errorMessage: string | null
}

function buildProductDetailEndpoint(productSlug: string) {
  if (CATALOG_PRODUCT_DETAIL_ENDPOINT.includes(':slug')) {
    return CATALOG_PRODUCT_DETAIL_ENDPOINT.replace(
      ':slug',
      encodeURIComponent(productSlug),
    )
  }

  return `${CATALOG_PRODUCT_DETAIL_ENDPOINT.replace(/\/$/, '')}/${encodeURIComponent(
    productSlug,
  )}`
}

export function useHerbalProductDetail(productSlug: string | undefined) {
  const [productDetailFetchState, setProductDetailFetchState] =
    useState<ProductDetailFetchState>({
      product: null,
      isLoading: true,
      errorMessage: null,
    })
  const [selectedOrderQuantity, setSelectedOrderQuantity] = useState(1)

  const loadProductDetail = useCallback(
    async (signal?: AbortSignal) => {
      await Promise.resolve()

      if (!productSlug) {
        setProductDetailFetchState({
          product: null,
          isLoading: false,
          errorMessage: null,
        })
        return
      }

      setProductDetailFetchState((currentState) => ({
        ...currentState,
        isLoading: true,
        errorMessage: null,
      }))

      try {
        const productDetailResponse = await fetch(buildProductDetailEndpoint(productSlug), {
          headers: { Accept: 'application/json' },
          signal,
        })

        if (!productDetailResponse.ok) {
          throw new Error(`Product detail API responded with ${productDetailResponse.status}`)
        }

        const productDetailPayload: unknown = await productDetailResponse.json()
        const productDetail = extractCatalogProductDetail(productDetailPayload)

        setProductDetailFetchState({
          product: productDetail?.is_active ? productDetail : null,
          isLoading: false,
          errorMessage: null,
        })
        setSelectedOrderQuantity(1)
      } catch (error) {
        if (error instanceof DOMException && error.name === 'AbortError') {
          return
        }

        setProductDetailFetchState({
          product: null,
          isLoading: false,
          errorMessage: friendlyProductDetailError,
        })
      }
    },
    [productSlug],
  )

  useEffect(() => {
    const productDetailRequestController = new AbortController()
    const initialProductDetailRequestId = window.setTimeout(() => {
      void loadProductDetail(productDetailRequestController.signal)
    }, 0)

    return () => {
      window.clearTimeout(initialProductDetailRequestId)
      productDetailRequestController.abort()
    }
  }, [loadProductDetail])

  const maximumOrderQuantity = productDetailFetchState.product?.stock_quantity ?? 1
  const canDecreaseOrderQuantity = selectedOrderQuantity > 1
  const canIncreaseOrderQuantity = selectedOrderQuantity < maximumOrderQuantity

  const productSubtotal = useMemo(() => {
    return (productDetailFetchState.product?.price ?? 0) * selectedOrderQuantity
  }, [productDetailFetchState.product?.price, selectedOrderQuantity])

  const decreaseOrderQuantity = useCallback(() => {
    setSelectedOrderQuantity((currentQuantity) => Math.max(1, currentQuantity - 1))
  }, [])

  const increaseOrderQuantity = useCallback(() => {
    setSelectedOrderQuantity((currentQuantity) =>
      Math.min(maximumOrderQuantity, currentQuantity + 1),
    )
  }, [maximumOrderQuantity])

  return {
    canDecreaseOrderQuantity,
    canIncreaseOrderQuantity,
    decreaseOrderQuantity,
    errorMessage: productDetailFetchState.errorMessage,
    increaseOrderQuantity,
    isLoading: productDetailFetchState.isLoading,
    maximumOrderQuantity,
    productDetail: productDetailFetchState.product,
    productSubtotal,
    retryProductDetailRequest: () => void loadProductDetail(),
    selectedOrderQuantity,
  }
}
