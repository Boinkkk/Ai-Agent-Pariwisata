import type { FormEvent } from 'react'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import {
  AlertCircle,
  CheckCircle2,
  Eye,
  EyeOff,
  Leaf,
  LockKeyhole,
  Mail,
  ShieldCheck,
  Sparkles,
} from 'lucide-react'
import { useCatalogAccountLogin } from './useCatalogAccountLogin'

function LoginBrandPanel() {
  return (
    <section className="relative hidden min-h-screen flex-1 overflow-hidden bg-emerald-950 px-12 py-10 text-white lg:flex lg:flex-col lg:justify-between">
      <div className="absolute inset-0 bg-[radial-gradient(circle_at_24%_18%,rgba(16,185,129,0.26),transparent_32%),radial-gradient(circle_at_76%_74%,rgba(245,158,11,0.2),transparent_30%)]" />
      <div className="relative z-10 inline-flex items-center gap-3">
        <span className="flex h-11 w-11 items-center justify-center rounded-md bg-white text-emerald-700">
          <Leaf aria-hidden="true" size={23} />
        </span>
        <span className="text-lg font-semibold tracking-wide">Herbali</span>
      </div>

      <div className="relative z-10 max-w-xl">
        <p className="mb-5 inline-flex items-center gap-2 rounded-full border border-white/15 bg-white/10 px-4 py-2 text-sm font-medium text-emerald-50">
          <Sparkles aria-hidden="true" size={16} />
          Portal katalog dan pesanan herbal
        </p>
        <h1 className="text-5xl font-semibold leading-tight tracking-normal">
          Masuk untuk mengelola katalog herbal dengan lebih tenang.
        </h1>
        <p className="mt-6 max-w-lg text-base leading-7 text-emerald-50/80">
          Akses dashboard produk, pantau pesanan, dan lanjutkan percakapan pelanggan
          dari satu ruang kerja yang rapi.
        </p>
      </div>

      <div className="relative z-10 grid grid-cols-2 gap-4">
        <div className="rounded-md border border-white/15 bg-white/10 p-5">
          <p className="text-3xl font-semibold">98%</p>
          <p className="mt-2 text-sm leading-6 text-emerald-50/75">
            Produk aktif siap tampil di katalog pelanggan.
          </p>
        </div>
        <div className="rounded-md border border-white/15 bg-white/10 p-5">
          <p className="text-3xl font-semibold">24/7</p>
          <p className="mt-2 text-sm leading-6 text-emerald-50/75">
            Chatbot membantu respons awal di luar jam operasional.
          </p>
        </div>
      </div>
    </section>
  )
}

interface LoginFieldProps {
  autoComplete: string
  errorMessage?: string
  icon: typeof Mail
  inputMode?: 'email'
  label: string
  name: string
  onChange: (value: string) => void
  placeholder: string
  type: 'email' | 'password' | 'text'
  value: string
  trailingControl?: React.ReactNode
}

function LoginField({
  autoComplete,
  errorMessage,
  icon: FieldIcon,
  inputMode,
  label,
  name,
  onChange,
  placeholder,
  type,
  value,
  trailingControl,
}: LoginFieldProps) {
  const fieldErrorId = `${name}-error`

  return (
    <div>
      <label htmlFor={name} className="text-sm font-semibold text-slate-800">
        {label}
      </label>
      <div className="relative mt-2">
        <FieldIcon
          aria-hidden="true"
          className="pointer-events-none absolute left-4 top-1/2 -translate-y-1/2 text-slate-400"
          size={18}
        />
        <input
          id={name}
          name={name}
          type={type}
          value={value}
          inputMode={inputMode}
          autoComplete={autoComplete}
          aria-invalid={errorMessage ? 'true' : 'false'}
          aria-describedby={errorMessage ? fieldErrorId : undefined}
          onChange={(event) => onChange(event.target.value)}
          placeholder={placeholder}
          className={`h-12 w-full rounded-md border bg-white pl-11 text-sm text-slate-950 transition placeholder:text-slate-400 focus:outline-none focus:ring-4 ${
            trailingControl ? 'pr-12' : 'pr-4'
          } ${
            errorMessage
              ? 'border-rose-300 focus:border-rose-500 focus:ring-rose-100'
              : 'border-slate-200 hover:border-slate-300 focus:border-emerald-500 focus:ring-emerald-100'
          }`}
        />
        {trailingControl ? (
          <div className="absolute right-2 top-1/2 -translate-y-1/2">
            {trailingControl}
          </div>
        ) : null}
      </div>
      {errorMessage ? (
        <p id={fieldErrorId} className="mt-2 text-sm font-medium text-rose-600">
          {errorMessage}
        </p>
      ) : null}
    </div>
  )
}

interface RememberedAccountPanelProps {
  hasRememberedCatalogAccount: boolean
  rememberedAccountEmail: string
}

function RememberedAccountPanel({
  hasRememberedCatalogAccount,
  rememberedAccountEmail,
}: RememberedAccountPanelProps) {
  if (hasRememberedCatalogAccount) {
    return (
      <div className="rounded-md border border-emerald-100 bg-emerald-50 px-4 py-3">
        <p className="inline-flex items-center gap-2 text-sm font-semibold text-emerald-800">
          <CheckCircle2 aria-hidden="true" size={16} />
          Akun tersimpan ditemukan
        </p>
        <p className="mt-1 text-sm text-emerald-700">{rememberedAccountEmail}</p>
      </div>
    )
  }

  return (
    <div className="rounded-md border border-dashed border-slate-200 bg-slate-50 px-4 py-3">
      <p className="text-sm font-semibold text-slate-800">Belum ada akun tersimpan</p>
      <p className="mt-1 text-sm leading-6 text-slate-500">
        Centang opsi ingat akun untuk mempercepat login berikutnya.
      </p>
    </div>
  )
}

export function LoginPage() {
  const navigate = useNavigate()
  const [isPasswordVisible, setIsPasswordVisible] = useState(false)
  const {
    catalogAccountCredentials,
    credentialFieldErrors,
    hasRememberedCatalogAccount,
    isLoginActionDisabled,
    isSubmittingLogin,
    loginErrorMessage,
    rememberedAccountEmail,
    submitCatalogAccountLogin,
    updateCatalogLoginEmail,
    updateCatalogLoginPassword,
    updateRememberCatalogAccountPreference,
  } = useCatalogAccountLogin({
    onAuthenticatedCatalogSession: () => navigate('/katalog'),
  })

  const submitLoginForm = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    void submitCatalogAccountLogin()
  }

  return (
    <main className="min-h-screen bg-slate-100 text-slate-950">
      <div className="mx-auto flex min-h-screen max-w-[1440px] bg-white">
        <LoginBrandPanel />

        <section className="flex min-h-screen w-full items-center justify-center px-5 py-10 sm:px-8 lg:w-[560px]">
          <div className="w-full max-w-md">
            <a
              href="/"
              className="mb-10 inline-flex items-center gap-3 rounded-md focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-emerald-600 lg:hidden"
              aria-label="Beranda Herbali"
            >
              <span className="flex h-10 w-10 items-center justify-center rounded-md bg-emerald-700 text-white">
                <Leaf aria-hidden="true" size={21} />
              </span>
              <span className="text-base font-semibold">Herbali</span>
            </a>

            <div className="mb-8">
              <p className="inline-flex items-center gap-2 text-sm font-semibold text-emerald-700">
                <ShieldCheck aria-hidden="true" size={17} />
                Akses aman dashboard
              </p>
              <h1 className="mt-3 text-3xl font-semibold tracking-normal text-slate-950">
                Masuk ke akun Anda
              </h1>
              <p className="mt-3 text-sm leading-6 text-slate-500">
                Gunakan email operasional untuk melanjutkan pengelolaan katalog dan
                pesanan.
              </p>
            </div>

            <RememberedAccountPanel
              hasRememberedCatalogAccount={hasRememberedCatalogAccount}
              rememberedAccountEmail={rememberedAccountEmail}
            />

            {loginErrorMessage ? (
              <div
                className="mt-5 flex gap-3 rounded-md border border-rose-200 bg-rose-50 px-4 py-3 text-sm text-rose-700"
                role="alert"
              >
                <AlertCircle aria-hidden="true" className="mt-0.5 shrink-0" size={17} />
                <p>{loginErrorMessage}</p>
              </div>
            ) : null}

            <form className="mt-6 space-y-5" onSubmit={submitLoginForm} noValidate>
              <LoginField
                autoComplete="email"
                errorMessage={credentialFieldErrors.emailAddress}
                icon={Mail}
                inputMode="email"
                label="Email"
                name="email-address"
                onChange={updateCatalogLoginEmail}
                placeholder="admin@herbali.id"
                type="email"
                value={catalogAccountCredentials.emailAddress}
              />

              <LoginField
                autoComplete="current-password"
                errorMessage={credentialFieldErrors.password}
                icon={LockKeyhole}
                label="Kata sandi"
                name="password"
                onChange={updateCatalogLoginPassword}
                placeholder="Masukkan kata sandi"
                type={isPasswordVisible ? 'text' : 'password'}
                value={catalogAccountCredentials.password}
                trailingControl={
                  <button
                    type="button"
                    onClick={() =>
                      setIsPasswordVisible((currentVisibility) => !currentVisibility)
                    }
                    className="inline-flex h-9 w-9 items-center justify-center rounded-md text-slate-500 transition hover:bg-slate-100 hover:text-slate-900 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-emerald-600"
                    aria-label={
                      isPasswordVisible
                        ? 'Sembunyikan kata sandi'
                        : 'Tampilkan kata sandi'
                    }
                    aria-pressed={isPasswordVisible}
                  >
                    {isPasswordVisible ? (
                      <EyeOff aria-hidden="true" size={18} />
                    ) : (
                      <Eye aria-hidden="true" size={18} />
                    )}
                  </button>
                }
              />

              <div className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
                <label className="inline-flex cursor-pointer items-center gap-3 text-sm font-medium text-slate-600">
                  <input
                    type="checkbox"
                    checked={catalogAccountCredentials.shouldRememberAccount}
                    onChange={(event) =>
                      updateRememberCatalogAccountPreference(event.target.checked)
                    }
                    className="h-4 w-4 rounded border-slate-300 text-emerald-600 focus:ring-4 focus:ring-emerald-100"
                  />
                  Ingat akun ini
                </label>
                <a
                  href="/forgot-password"
                  className="rounded-md text-sm font-semibold text-emerald-700 transition hover:text-emerald-800 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-emerald-600"
                >
                  Lupa kata sandi?
                </a>
              </div>

              <button
                type="submit"
                disabled={isLoginActionDisabled}
                className="inline-flex h-12 w-full items-center justify-center rounded-md bg-emerald-700 px-4 text-sm font-semibold text-white transition hover:bg-emerald-800 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-emerald-700 disabled:cursor-not-allowed disabled:bg-slate-300"
              >
                {isSubmittingLogin ? (
                  <span className="inline-flex items-center gap-2">
                    <span
                      className="h-4 w-4 animate-spin rounded-full border-2 border-white/40 border-t-white"
                      aria-hidden="true"
                    />
                    Memproses login
                  </span>
                ) : (
                  'Masuk'
                )}
              </button>
            </form>
          </div>
        </section>
      </div>
    </main>
  )
}
