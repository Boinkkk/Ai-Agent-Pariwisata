import {
  BotMessageSquare,
  BriefcaseBusiness,
  HeartPulse,
  LogOut,
  Menu,
  PackageSearch,
  Search,
  ShoppingBag,
  SlidersHorizontal,
  Sparkles,
  Star,
  User,
} from 'lucide-react'
import { useState } from 'react'
import type {
  CatalogSortPreference,
  HerbalCatalogProduct,
} from './catalogTypes'
import { useHerbalProductCatalog } from './useHerbalProductCatalog'

const rupiahFormatter = new Intl.NumberFormat('id-ID', {
  currency: 'IDR',
  maximumFractionDigits: 0,
  style: 'currency',
})

const catalogNavigationItems = [
  { label: 'Katalog', icon: HeartPulse, isActive: true },
  { label: 'Order', icon: BriefcaseBusiness, isActive: false },
  { label: 'User Profile', icon: User, isActive: false },
  { label: 'Chatbot', icon: BotMessageSquare, isActive: false },
]

interface CatalogToolbarProps {
  availableRegencies: string[]
  catalogSearchTerm: string
  selectedRegency: string
  sortPreference: CatalogSortPreference
  onRegencyChange: (regency: string) => void
  onSearchTermChange: (searchTerm: string) => void
  onSortPreferenceChange: (sortPreference: CatalogSortPreference) => void
}

interface CatalogGridStatusProps {
  onRetryCatalogRequest?: () => void
  statusMessage: string
  statusTitle: string
  tone: 'empty' | 'error'
}

function CatalogSidebar() {
  return (
    <aside className="flex w-full flex-col border-r border-slate-200 bg-white md:min-h-screen md:w-56">
      <div className="flex h-18 items-center justify-between border-b border-slate-100 px-6">
        <a
          href="/"
          className="inline-flex items-center gap-2 rounded-md focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-emerald-600"
          aria-label="Beranda katalog herbal"
        >
          <span className="flex h-10 w-10 items-center justify-center rounded-full bg-emerald-50 text-emerald-600">
            <Sparkles aria-hidden="true" size={22} />
          </span>
          <span className="text-sm font-semibold text-slate-900">Herbali</span>
        </a>
        <button
          type="button"
          className="inline-flex h-9 w-9 items-center justify-center rounded-md text-slate-500 transition hover:bg-slate-100 hover:text-slate-900 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-emerald-600 md:hidden"
          aria-label="Buka navigasi katalog"
        >
          <Menu aria-hidden="true" size={18} />
        </button>
      </div>

      <nav className="flex flex-1 flex-row gap-2 overflow-x-auto px-4 py-5 md:flex-col md:overflow-visible">
        {catalogNavigationItems.map((navigationItem) => {
          const NavigationIcon = navigationItem.icon

          return (
            <a
              key={navigationItem.label}
              href={`#${navigationItem.label.toLowerCase().replace(/\s+/g, '-')}`}
              aria-current={navigationItem.isActive ? 'page' : undefined}
              className={`inline-flex min-w-max items-center gap-3 rounded-md px-4 py-3 text-sm font-medium transition focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-orange-500 ${
                navigationItem.isActive
                  ? 'bg-orange-100 text-orange-700'
                  : 'text-slate-500 hover:bg-slate-50 hover:text-slate-900'
              }`}
            >
              <NavigationIcon aria-hidden="true" size={18} />
              {navigationItem.label}
            </a>
          )
        })}
      </nav>

      <div className="hidden border-t border-slate-100 px-4 py-5 md:block">
        <button
          type="button"
          className="inline-flex w-full items-center gap-3 rounded-md px-4 py-3 text-sm font-medium text-orange-600 transition hover:bg-orange-50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-orange-500"
        >
          <LogOut aria-hidden="true" size={18} />
          Logout
        </button>
      </div>
    </aside>
  )
}

function CatalogToolbar({
  availableRegencies,
  catalogSearchTerm,
  selectedRegency,
  sortPreference,
  onRegencyChange,
  onSearchTermChange,
  onSortPreferenceChange,
}: CatalogToolbarProps) {
  return (
    <div className="flex flex-col gap-4 border-b border-slate-200 bg-white px-5 py-4 lg:h-18 lg:flex-row lg:items-center lg:px-8">
      <div className="relative max-w-2xl flex-1">
        <label htmlFor="catalog-search" className="sr-only">
          Cari jamu, manfaat, atau komposisi
        </label>
        <Search
          aria-hidden="true"
          className="pointer-events-none absolute left-4 top-1/2 -translate-y-1/2 text-slate-400"
          size={18}
        />
        <input
          id="catalog-search"
          type="search"
          value={catalogSearchTerm}
          onChange={(event) => onSearchTermChange(event.target.value)}
          placeholder="Search something..."
          className="h-11 w-full rounded-md border border-slate-200 bg-white pl-11 pr-4 text-sm text-slate-900 transition placeholder:text-slate-300 hover:border-slate-300 focus:border-emerald-500 focus:outline-none focus:ring-4 focus:ring-emerald-100"
        />
      </div>

      <div className="flex flex-wrap items-center gap-3">
        <span className="inline-flex items-center gap-2 text-xs font-semibold text-slate-500">
          <SlidersHorizontal aria-hidden="true" size={16} />
          Filter by
        </span>

        <label htmlFor="catalog-regency-filter" className="sr-only">
          Filter katalog berdasarkan daerah produksi
        </label>
        <select
          id="catalog-regency-filter"
          value={selectedRegency}
          onChange={(event) => onRegencyChange(event.target.value)}
          className="h-9 rounded-md border border-slate-200 bg-white px-3 text-xs font-medium text-slate-600 transition hover:border-slate-300 focus:border-emerald-500 focus:outline-none focus:ring-4 focus:ring-emerald-100"
        >
          <option value="all">Semua daerah</option>
          {availableRegencies.map((regency) => (
            <option key={regency} value={regency}>
              {regency}
            </option>
          ))}
        </select>

        <label htmlFor="catalog-sort-filter" className="sr-only">
          Urutkan produk katalog
        </label>
        <select
          id="catalog-sort-filter"
          value={sortPreference}
          onChange={(event) =>
            onSortPreferenceChange(event.target.value as CatalogSortPreference)
          }
          className="h-9 rounded-md border border-slate-200 bg-white px-3 text-xs font-medium text-slate-600 transition hover:border-slate-300 focus:border-emerald-500 focus:outline-none focus:ring-4 focus:ring-emerald-100"
        >
          <option value="featured">Rekomendasi</option>
          <option value="price-ascending">Harga termurah</option>
          <option value="rating-descending">Rating tertinggi</option>
        </select>
      </div>
    </div>
  )
}

function CatalogProductImage({ product }: { product: HerbalCatalogProduct }) {
  const [hasImageLoadFailed, setHasImageLoadFailed] = useState(false)
  const shouldShowProductImage = product.image_url && !hasImageLoadFailed

  return (
    <div className="relative aspect-square overflow-hidden rounded-md bg-slate-200">
      {shouldShowProductImage ? (
        <img
          src={product.image_url}
          alt={product.name}
          className="h-full w-full object-cover transition duration-300 group-hover:scale-105"
          loading="lazy"
          onError={() => setHasImageLoadFailed(true)}
        />
      ) : (
        <div className="flex h-full w-full items-center justify-center bg-emerald-50 text-emerald-600">
          <PackageSearch aria-hidden="true" size={38} />
        </div>
      )}
    </div>
  )
}

function CatalogProductCard({ product }: { product: HerbalCatalogProduct }) {
  const productDetailUrl = `/katalog/${product.slug}`

  return (
    <article className="group">
      <a
        href={productDetailUrl}
        className="block rounded-md focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-emerald-600"
        aria-label={`Lihat detail ${product.name}`}
      >
        <CatalogProductImage product={product} />
        <div className="mt-3 space-y-2 text-left">
          <div className="flex items-start justify-between gap-3">
            <h2 className="line-clamp-2 text-sm font-semibold leading-5 text-slate-950">
              {product.name}
            </h2>
            <span className="inline-flex items-center gap-1 text-xs font-semibold text-amber-600">
              <Star aria-hidden="true" size={14} fill="currentColor" />
              {product.average_rating.toFixed(1)}
            </span>
          </div>
          <p className="line-clamp-2 text-xs leading-5 text-slate-500">
            {product.description}
          </p>
          <div className="flex flex-wrap items-center gap-2 text-xs text-slate-500">
            <span className="font-semibold text-emerald-700">
              {rupiahFormatter.format(product.price)}
            </span>
            <span aria-hidden="true">/</span>
            <span>{product.weight_grams} gr</span>
          </div>
          <p className="text-xs font-medium text-slate-400">
            Stok {product.stock_quantity} - {product.regency || product.production_location}
          </p>
        </div>
      </a>
    </article>
  )
}

function CatalogSkeletonGrid() {
  return (
    <div
      className="grid grid-cols-2 gap-x-7 gap-y-8 sm:grid-cols-3 xl:grid-cols-5"
      aria-label="Memuat produk katalog"
    >
      {Array.from({ length: 10 }).map((_, skeletonIndex) => (
        <div key={skeletonIndex} className="animate-pulse">
          <div className="aspect-square rounded-md bg-slate-200" />
          <div className="mt-3 h-3 w-11/12 rounded bg-slate-200" />
          <div className="mt-2 h-3 w-8/12 rounded bg-slate-200" />
        </div>
      ))}
    </div>
  )
}

function CatalogGridStatus({
  onRetryCatalogRequest,
  statusMessage,
  statusTitle,
  tone,
}: CatalogGridStatusProps) {
  const isError = tone === 'error'

  return (
    <div className="flex min-h-96 items-center justify-center rounded-md border border-dashed border-slate-200 bg-white px-6 py-12 text-center">
      <div className="max-w-md">
        <div
          className={`mx-auto flex h-14 w-14 items-center justify-center rounded-full ${
            isError ? 'bg-rose-50 text-rose-600' : 'bg-emerald-50 text-emerald-600'
          }`}
        >
          <PackageSearch aria-hidden="true" size={26} />
        </div>
        <h2 className="mt-5 text-lg font-semibold text-slate-950">{statusTitle}</h2>
        <p className="mt-2 text-sm leading-6 text-slate-500">{statusMessage}</p>
        {onRetryCatalogRequest ? (
          <button
            type="button"
            onClick={onRetryCatalogRequest}
            className="mt-6 inline-flex h-10 items-center justify-center rounded-md bg-emerald-600 px-4 text-sm font-semibold text-white transition hover:bg-emerald-700 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-emerald-600 disabled:cursor-not-allowed disabled:bg-slate-300"
          >
            Coba lagi
          </button>
        ) : null}
      </div>
    </div>
  )
}

export function CatalogPage() {
  const {
    availableRegencies,
    catalogSearchTerm,
    errorMessage,
    isLoading,
    productsTotal,
    retryCatalogRequest,
    selectedRegency,
    setCatalogSearchTerm,
    setSelectedRegency,
    setSortPreference,
    sortPreference,
    visibleCatalogProducts,
  } = useHerbalProductCatalog()

  const visibleProductsCount = visibleCatalogProducts.length
  const catalogRangeSummary =
    visibleProductsCount > 0
      ? `Showing 1-${visibleProductsCount} of ${productsTotal} result`
      : 'Showing 0 result'

  return (
    <div className="min-h-screen bg-slate-100 text-slate-900">
      <div className="mx-auto flex min-h-screen max-w-[1440px] flex-col border-x border-slate-200 bg-white md:flex-row">
        <CatalogSidebar />

        <main className="min-w-0 flex-1 bg-slate-50">
          <CatalogToolbar
            availableRegencies={availableRegencies}
            catalogSearchTerm={catalogSearchTerm}
            selectedRegency={selectedRegency}
            sortPreference={sortPreference}
            onRegencyChange={setSelectedRegency}
            onSearchTermChange={setCatalogSearchTerm}
            onSortPreferenceChange={setSortPreference}
          />

          <section className="px-5 py-7 lg:px-12" aria-labelledby="catalog-heading">
            <div className="mb-6 flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
              <div>
                <p className="text-xs font-medium text-slate-500">{catalogRangeSummary}</p>
                <h1 id="catalog-heading" className="sr-only">
                  Katalog produk jamu dan herbal
                </h1>
              </div>
              <button
                type="button"
                className="inline-flex h-10 items-center justify-center gap-2 rounded-md border border-slate-200 bg-white px-4 text-sm font-semibold text-slate-700 transition hover:border-slate-300 hover:bg-slate-50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-emerald-600"
              >
                <ShoppingBag aria-hidden="true" size={17} />
                Keranjang
              </button>
            </div>

            {isLoading ? <CatalogSkeletonGrid /> : null}

            {!isLoading && errorMessage ? (
              <CatalogGridStatus
                statusTitle="Katalog belum tersedia"
                statusMessage={errorMessage}
                tone="error"
                onRetryCatalogRequest={retryCatalogRequest}
              />
            ) : null}

            {!isLoading && !errorMessage && visibleProductsCount === 0 ? (
              <CatalogGridStatus
                statusTitle="Produk tidak ditemukan"
                statusMessage="Coba gunakan kata kunci lain atau ubah filter daerah produksi."
                tone="empty"
              />
            ) : null}

            {!isLoading && !errorMessage && visibleProductsCount > 0 ? (
              <div className="grid grid-cols-2 gap-x-7 gap-y-8 sm:grid-cols-3 xl:grid-cols-5">
                {visibleCatalogProducts.map((product) => (
                  <CatalogProductCard key={product.id} product={product} />
                ))}
              </div>
            ) : null}
          </section>
        </main>
      </div>
    </div>
  )
}
