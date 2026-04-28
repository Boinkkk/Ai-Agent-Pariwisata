import {
  ArrowLeft,
  BotMessageSquare,
  BriefcaseBusiness,
  ClipboardList,
  Factory,
  HeartPulse,
  Leaf,
  LogOut,
  MapPin,
  Menu,
  Minus,
  PackageSearch,
  Plus,
  Search,
  ShieldCheck,
  ShoppingBag,
  Star,
  User,
  UserCircle,
} from 'lucide-react'
import { useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import type { HerbalCatalogProduct } from './catalogTypes'
import { useHerbalProductDetail } from './useHerbalProductDetail'

const detailRupiahFormatter = new Intl.NumberFormat('id-ID', {
  currency: 'IDR',
  maximumFractionDigits: 0,
  style: 'currency',
})

const productDetailNavigationItems = [
  { label: 'Katalog', icon: HeartPulse, isActive: true, to: '/katalog' },
  { label: 'Order', icon: BriefcaseBusiness, isActive: false, to: '#order' },
  { label: 'User Profile', icon: User, isActive: false, to: '#user-profile' },
  { label: 'Chatbot', icon: BotMessageSquare, isActive: false, to: '#chatbot' },
]

function ProductDetailSidebar() {
  return (
    <aside className="flex w-full flex-col border-r border-slate-200 bg-white md:min-h-screen md:w-56">
      <div className="flex h-18 items-center justify-between border-b border-slate-100 px-6">
        <Link
          to="/katalog"
          className="inline-flex items-center gap-2 rounded-md focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-emerald-600"
          aria-label="Kembali ke katalog Herbali"
        >
          <span className="flex h-10 w-10 items-center justify-center rounded-full bg-emerald-50 text-emerald-600">
            <Leaf aria-hidden="true" size={22} />
          </span>
          <span className="text-sm font-semibold text-slate-900">Herbali</span>
        </Link>
        <button
          type="button"
          className="inline-flex h-9 w-9 items-center justify-center rounded-md text-slate-500 transition hover:bg-slate-100 hover:text-slate-900 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-emerald-600 md:hidden"
          aria-label="Buka navigasi detail produk"
        >
          <Menu aria-hidden="true" size={18} />
        </button>
      </div>

      <nav className="flex flex-1 flex-row gap-2 overflow-x-auto px-4 py-5 md:flex-col md:overflow-visible">
        {productDetailNavigationItems.map((navigationItem) => {
          const NavigationIcon = navigationItem.icon

          return (
            <Link
              key={navigationItem.label}
              to={navigationItem.to}
              aria-current={navigationItem.isActive ? 'page' : undefined}
              className={`inline-flex min-w-max items-center gap-3 rounded-md px-4 py-3 text-sm font-medium transition focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-orange-500 ${
                navigationItem.isActive
                  ? 'bg-orange-100 text-orange-700'
                  : 'text-slate-500 hover:bg-slate-50 hover:text-slate-900'
              }`}
            >
              <NavigationIcon aria-hidden="true" size={18} />
              {navigationItem.label}
            </Link>
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

function ProductDetailTopbar() {
  return (
    <header className="flex flex-col gap-4 border-b border-slate-200 bg-white px-5 py-4 lg:h-18 lg:flex-row lg:items-center lg:justify-between lg:px-8">
      <div className="relative max-w-xl flex-1">
        <label htmlFor="product-detail-search" className="sr-only">
          Cari produk dari halaman detail
        </label>
        <Search
          aria-hidden="true"
          className="pointer-events-none absolute left-4 top-1/2 -translate-y-1/2 text-slate-400"
          size={18}
        />
        <input
          id="product-detail-search"
          type="search"
          placeholder="Search something..."
          className="h-11 w-full rounded-md border border-slate-200 bg-white pl-11 pr-4 text-sm text-slate-900 transition placeholder:text-slate-300 hover:border-slate-300 focus:border-emerald-500 focus:outline-none focus:ring-4 focus:ring-emerald-100"
        />
      </div>

      <div className="flex items-center gap-3">
        <button
          type="button"
          className="inline-flex h-10 w-10 items-center justify-center rounded-md border border-slate-200 bg-white text-slate-600 transition hover:border-slate-300 hover:bg-slate-50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-emerald-600"
          aria-label="Buka keranjang"
        >
          <ShoppingBag aria-hidden="true" size={18} />
        </button>
        <button
          type="button"
          className="inline-flex h-10 w-10 items-center justify-center rounded-full border border-slate-200 bg-white text-slate-600 transition hover:border-slate-300 hover:bg-slate-50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-emerald-600"
          aria-label="Buka profil pengguna"
        >
          <UserCircle aria-hidden="true" size={22} />
        </button>
      </div>
    </header>
  )
}

function ProductDetailImage({ product }: { product: HerbalCatalogProduct }) {
  const [hasImageLoadFailed, setHasImageLoadFailed] = useState(false)
  const shouldShowProductImage = product.image_url && !hasImageLoadFailed

  return (
    <div className="aspect-square w-full overflow-hidden rounded-md bg-slate-200 sm:max-w-[220px]">
      {shouldShowProductImage ? (
        <img
          src={product.image_url}
          alt={product.name}
          className="h-full w-full object-cover"
          onError={() => setHasImageLoadFailed(true)}
        />
      ) : (
        <div className="flex h-full w-full items-center justify-center bg-emerald-50 text-emerald-600">
          <PackageSearch aria-hidden="true" size={44} />
        </div>
      )}
    </div>
  )
}

function ProductAttributeRow({
  label,
  value,
}: {
  label: string
  value: string | number
}) {
  return (
    <div className="flex items-start justify-between gap-4 border-b border-slate-100 py-3 last:border-b-0">
      <dt className="text-sm font-medium text-slate-500">{label}</dt>
      <dd className="text-right text-sm font-semibold text-slate-900">{value || '-'}</dd>
    </div>
  )
}

function ProductInformationSection({ product }: { product: HerbalCatalogProduct }) {
  return (
    <div className="mt-7 grid gap-6 lg:grid-cols-2">
      <section aria-labelledby="product-benefit-heading">
        <h2
          id="product-benefit-heading"
          className="text-base font-semibold text-slate-950"
        >
          Deskripsi Produk
        </h2>
        <p className="mt-2 text-sm leading-6 text-slate-600">{product.description}</p>

        <h3 className="mt-5 text-base font-semibold text-slate-950">Aturan Pakai</h3>
        <p className="mt-2 text-sm leading-6 text-slate-600">
          {product.directions || 'Instruksi penggunaan belum tersedia.'}
        </p>

        <h3 className="mt-5 text-base font-semibold text-slate-950">Manfaat</h3>
        <p className="mt-2 text-sm leading-6 text-slate-600">
          {product.benefit || 'Manfaat produk belum tersedia.'}
        </p>
      </section>

      <section aria-labelledby="product-identity-heading">
        <h2
          id="product-identity-heading"
          className="text-base font-semibold text-slate-950"
        >
          Informasi Produk
        </h2>
        <dl className="mt-2 rounded-md border border-slate-100 bg-white px-4">
          <ProductAttributeRow label="Komposisi" value={product.composition} />
          <ProductAttributeRow label="Produsen" value={product.manufacturer} />
          <ProductAttributeRow label="Lokasi produksi" value={product.production_location} />
          <ProductAttributeRow label="Wilayah" value={product.regency} />
          <ProductAttributeRow
            label="Lisensi"
            value={`${product.licensing || '-'} ${product.licensing_number || ''}`}
          />
        </dl>
      </section>
    </div>
  )
}

interface ProductPurchasePanelProps {
  canDecreaseOrderQuantity: boolean
  canIncreaseOrderQuantity: boolean
  onDecreaseOrderQuantity: () => void
  onIncreaseOrderQuantity: () => void
  product: HerbalCatalogProduct
  productSubtotal: number
  selectedOrderQuantity: number
}

function ProductPurchasePanel({
  canDecreaseOrderQuantity,
  canIncreaseOrderQuantity,
  onDecreaseOrderQuantity,
  onIncreaseOrderQuantity,
  product,
  productSubtotal,
  selectedOrderQuantity,
}: ProductPurchasePanelProps) {
  const isOutOfStock = product.stock_quantity < 1

  return (
    <aside className="rounded-md border border-slate-200 bg-white p-5 shadow-sm lg:w-72">
      <h2 className="text-base font-semibold text-slate-950">Atur Jumlah</h2>

      <div className="mt-4 flex gap-3">
        <div className="h-16 w-16 shrink-0 overflow-hidden rounded-md bg-slate-100">
          {product.image_url ? (
            <img src={product.image_url} alt="" className="h-full w-full object-cover" />
          ) : (
            <div className="flex h-full w-full items-center justify-center text-emerald-600">
              <Leaf aria-hidden="true" size={22} />
            </div>
          )}
        </div>
        <div className="min-w-0">
          <p className="truncate text-sm font-semibold text-slate-950">{product.name}</p>
          <p className="mt-1 text-xs text-slate-500">{product.weight_grams} gr</p>
        </div>
      </div>

      <div className="mt-5 flex items-center justify-between gap-3">
        <div
          className="inline-flex h-10 items-center rounded-md border border-slate-200"
          aria-label="Kontrol jumlah produk"
        >
          <button
            type="button"
            onClick={onDecreaseOrderQuantity}
            disabled={!canDecreaseOrderQuantity || isOutOfStock}
            className="inline-flex h-10 w-10 items-center justify-center text-orange-600 transition hover:bg-orange-50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-orange-500 disabled:cursor-not-allowed disabled:text-slate-300"
            aria-label="Kurangi jumlah produk"
          >
            <Minus aria-hidden="true" size={16} />
          </button>
          <output
            className="w-9 text-center text-sm font-semibold text-slate-950"
            aria-live="polite"
          >
            {isOutOfStock ? 0 : selectedOrderQuantity}
          </output>
          <button
            type="button"
            onClick={onIncreaseOrderQuantity}
            disabled={!canIncreaseOrderQuantity || isOutOfStock}
            className="inline-flex h-10 w-10 items-center justify-center text-orange-600 transition hover:bg-orange-50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-orange-500 disabled:cursor-not-allowed disabled:text-slate-300"
            aria-label="Tambah jumlah produk"
          >
            <Plus aria-hidden="true" size={16} />
          </button>
        </div>
        <p className="text-xs font-medium text-slate-500">Stok {product.stock_quantity}</p>
      </div>

      <div className="mt-5 flex items-center justify-between border-t border-slate-100 pt-4">
        <span className="text-xs font-medium text-slate-500">Subtotal</span>
        <strong className="text-sm font-semibold text-orange-600">
          {detailRupiahFormatter.format(isOutOfStock ? 0 : productSubtotal)}
        </strong>
      </div>

      <div className="mt-8 grid grid-cols-2 gap-3">
        <button
          type="button"
          disabled={isOutOfStock}
          className="inline-flex h-10 items-center justify-center rounded-md border border-orange-500 bg-white px-3 text-xs font-semibold text-orange-600 transition hover:bg-orange-50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-orange-500 disabled:cursor-not-allowed disabled:border-slate-200 disabled:text-slate-300"
        >
          Beli langsung
        </button>
        <button
          type="button"
          disabled={isOutOfStock}
          className="inline-flex h-10 items-center justify-center rounded-md bg-orange-500 px-3 text-xs font-semibold text-white transition hover:bg-orange-600 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-orange-500 disabled:cursor-not-allowed disabled:bg-slate-300"
        >
          + Keranjang
        </button>
      </div>
    </aside>
  )
}

function ProductDetailSkeleton() {
  return (
    <div className="animate-pulse">
      <div className="flex flex-col gap-7 lg:flex-row lg:items-start lg:justify-between">
        <div className="flex flex-1 flex-col gap-6 sm:flex-row">
          <div className="aspect-square w-full rounded-md bg-slate-200 sm:max-w-[220px]" />
          <div className="flex-1 space-y-4">
            <div className="h-7 w-52 rounded bg-slate-200" />
            <div className="h-4 w-40 rounded bg-slate-200" />
            <div className="h-8 w-36 rounded bg-slate-200" />
            <div className="h-px w-full bg-slate-200" />
            <div className="h-4 w-32 rounded bg-slate-200" />
            <div className="h-16 w-full rounded bg-slate-200" />
          </div>
        </div>
        <div className="h-72 rounded-md bg-slate-200 lg:w-72" />
      </div>
    </div>
  )
}

function ProductDetailStatus({
  onRetryProductDetailRequest,
  statusMessage,
  statusTitle,
  tone,
}: {
  onRetryProductDetailRequest?: () => void
  statusMessage: string
  statusTitle: string
  tone: 'empty' | 'error'
}) {
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
        {onRetryProductDetailRequest ? (
          <button
            type="button"
            onClick={onRetryProductDetailRequest}
            className="mt-6 inline-flex h-10 items-center justify-center rounded-md bg-emerald-600 px-4 text-sm font-semibold text-white transition hover:bg-emerald-700 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-emerald-600"
          >
            Coba lagi
          </button>
        ) : null}
      </div>
    </div>
  )
}

function ProductDetailContent({
  canDecreaseOrderQuantity,
  canIncreaseOrderQuantity,
  decreaseOrderQuantity,
  increaseOrderQuantity,
  product,
  productSubtotal,
  selectedOrderQuantity,
}: {
  canDecreaseOrderQuantity: boolean
  canIncreaseOrderQuantity: boolean
  decreaseOrderQuantity: () => void
  increaseOrderQuantity: () => void
  product: HerbalCatalogProduct
  productSubtotal: number
  selectedOrderQuantity: number
}) {
  return (
    <div className="flex flex-col gap-7 lg:flex-row lg:items-start lg:justify-between">
      <div className="min-w-0 flex-1">
        <Link
          to="/katalog"
          className="mb-6 inline-flex items-center gap-2 rounded-md text-sm font-semibold text-slate-500 transition hover:text-slate-900 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-emerald-600"
        >
          <ArrowLeft aria-hidden="true" size={17} />
          Kembali ke katalog
        </Link>

        <div className="flex flex-col gap-6 sm:flex-row">
          <ProductDetailImage product={product} />
          <section className="min-w-0 flex-1" aria-labelledby="product-detail-heading">
            <div className="flex flex-wrap items-center gap-3">
              <span className="inline-flex items-center gap-1 rounded-full bg-emerald-50 px-3 py-1 text-xs font-semibold text-emerald-700">
                <ShieldCheck aria-hidden="true" size={14} />
                {product.licensing || 'Terverifikasi'}
              </span>
              <span className="inline-flex items-center gap-1 rounded-full bg-amber-50 px-3 py-1 text-xs font-semibold text-amber-700">
                <Star aria-hidden="true" size={14} fill="currentColor" />
                {product.average_rating.toFixed(1)}
              </span>
            </div>

            <h1
              id="product-detail-heading"
              className="mt-4 text-2xl font-semibold tracking-normal text-slate-950"
            >
              {product.name}
            </h1>
            <p className="mt-2 text-xs font-medium text-slate-500">
              Terjual 200+ pcs - {product.average_rating.toFixed(1)} ({product.weight_grams}
              gr)
            </p>
            <p className="mt-4 text-3xl font-semibold text-orange-500">
              {detailRupiahFormatter.format(product.price)}
            </p>

            <div className="mt-6 grid gap-3 border-y border-slate-100 py-4 sm:grid-cols-2">
              <p className="inline-flex items-center gap-2 text-sm text-slate-600">
                <Factory aria-hidden="true" size={17} className="text-slate-400" />
                {product.manufacturer}
              </p>
              <p className="inline-flex items-center gap-2 text-sm text-slate-600">
                <MapPin aria-hidden="true" size={17} className="text-slate-400" />
                {product.production_location || product.regency}
              </p>
              <p className="inline-flex items-center gap-2 text-sm text-slate-600">
                <ClipboardList aria-hidden="true" size={17} className="text-slate-400" />
                {product.composition}
              </p>
            </div>
          </section>
        </div>

        <ProductInformationSection product={product} />
      </div>

      <ProductPurchasePanel
        canDecreaseOrderQuantity={canDecreaseOrderQuantity}
        canIncreaseOrderQuantity={canIncreaseOrderQuantity}
        onDecreaseOrderQuantity={decreaseOrderQuantity}
        onIncreaseOrderQuantity={increaseOrderQuantity}
        product={product}
        productSubtotal={productSubtotal}
        selectedOrderQuantity={selectedOrderQuantity}
      />
    </div>
  )
}

export function ProductDetailPage() {
  const { productSlug } = useParams<{ productSlug: string }>()
  const {
    canDecreaseOrderQuantity,
    canIncreaseOrderQuantity,
    decreaseOrderQuantity,
    errorMessage,
    increaseOrderQuantity,
    isLoading,
    productDetail,
    productSubtotal,
    retryProductDetailRequest,
    selectedOrderQuantity,
  } = useHerbalProductDetail(productSlug)

  return (
    <div className="min-h-screen bg-slate-100 text-slate-900">
      <div className="mx-auto flex min-h-screen max-w-[1440px] flex-col border-x border-slate-200 bg-white md:flex-row">
        <ProductDetailSidebar />

        <main className="min-w-0 flex-1 bg-slate-50">
          <ProductDetailTopbar />

          <section className="px-5 py-7 lg:px-10" aria-live="polite">
            {isLoading ? <ProductDetailSkeleton /> : null}

            {!isLoading && errorMessage ? (
              <ProductDetailStatus
                statusTitle="Detail produk belum tersedia"
                statusMessage={errorMessage}
                tone="error"
                onRetryProductDetailRequest={retryProductDetailRequest}
              />
            ) : null}

            {!isLoading && !errorMessage && !productDetail ? (
              <ProductDetailStatus
                statusTitle="Produk tidak ditemukan"
                statusMessage="Produk ini mungkin tidak aktif atau tautan katalog sudah berubah."
                tone="empty"
              />
            ) : null}

            {!isLoading && !errorMessage && productDetail ? (
              <ProductDetailContent
                canDecreaseOrderQuantity={canDecreaseOrderQuantity}
                canIncreaseOrderQuantity={canIncreaseOrderQuantity}
                decreaseOrderQuantity={decreaseOrderQuantity}
                increaseOrderQuantity={increaseOrderQuantity}
                product={productDetail}
                productSubtotal={productSubtotal}
                selectedOrderQuantity={selectedOrderQuantity}
              />
            ) : null}
          </section>
        </main>
      </div>
    </div>
  )
}
