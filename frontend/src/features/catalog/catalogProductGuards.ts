import type { HerbalCatalogProduct } from './catalogTypes'

export function isHerbalCatalogProduct(
  candidate: unknown,
): candidate is HerbalCatalogProduct {
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

export function extractCatalogProducts(apiPayload: unknown): HerbalCatalogProduct[] {
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

export function extractCatalogProductDetail(
  apiPayload: unknown,
): HerbalCatalogProduct | null {
  if (isHerbalCatalogProduct(apiPayload)) {
    return apiPayload
  }

  if (!apiPayload || typeof apiPayload !== 'object') {
    return null
  }

  const possiblePayload = apiPayload as {
    data?: unknown
    product?: unknown
    item?: unknown
  }

  const possibleProduct =
    possiblePayload.data ?? possiblePayload.product ?? possiblePayload.item

  return isHerbalCatalogProduct(possibleProduct) ? possibleProduct : null
}
