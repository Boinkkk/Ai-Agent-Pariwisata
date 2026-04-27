export interface HerbalCatalogProduct {
  id: string
  category_id: number
  name: string
  slug: string
  description: string
  price: number
  stock_quantity: number
  weight_grams: number
  image_url: string
  is_active: boolean
  average_rating: number
  benefit: string
  composition: string
  directions: string
  storage_instructions: string
  manufacturer: string
  marketing_location: string
  production_location: string
  regency: string
  licensing: string
  licensing_number: string
  created_at: string
  updated_at: string
}

export type CatalogSortPreference = 'featured' | 'price-ascending' | 'rating-descending'

export interface CatalogFetchState {
  products: HerbalCatalogProduct[]
  isLoading: boolean
  errorMessage: string | null
}
