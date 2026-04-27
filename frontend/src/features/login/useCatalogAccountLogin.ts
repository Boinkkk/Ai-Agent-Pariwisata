import { useCallback, useState } from 'react'
import type {
  AuthenticatedCatalogSession,
  CatalogAccountCredentials,
  LoginApiResponse,
  LoginFieldErrors,
} from './loginTypes'

const ACCOUNT_LOGIN_ENDPOINT =
  import.meta.env.VITE_ACCOUNT_LOGIN_ENDPOINT ?? '/api/auth/login'
const rememberedCatalogAccountKey = 'herbali.remembered-account-email'
const authenticatedCatalogSessionKey = 'herbali.catalog-session'

const friendlyLoginErrorMessage =
  'Login belum berhasil. Periksa email dan kata sandi, lalu coba lagi.'

interface UseCatalogAccountLoginOptions {
  onAuthenticatedCatalogSession?: (session: AuthenticatedCatalogSession) => void
}

function readRememberedCatalogAccount() {
  try {
    return window.localStorage.getItem(rememberedCatalogAccountKey) ?? ''
  } catch {
    return ''
  }
}

function persistRememberedCatalogAccount(credentials: CatalogAccountCredentials) {
  try {
    if (credentials.shouldRememberAccount) {
      window.localStorage.setItem(
        rememberedCatalogAccountKey,
        credentials.emailAddress.trim(),
      )
      return
    }

    window.localStorage.removeItem(rememberedCatalogAccountKey)
  } catch {
    // Browser privacy settings can block storage; login should still continue.
  }
}

function persistAuthenticatedCatalogSession(session: AuthenticatedCatalogSession) {
  try {
    window.localStorage.setItem(authenticatedCatalogSessionKey, JSON.stringify(session))
  } catch {
    // Session persistence is progressive enhancement for this frontend.
  }
}

function validateCatalogAccountCredentials(
  credentials: CatalogAccountCredentials,
): LoginFieldErrors {
  const nextFieldErrors: LoginFieldErrors = {}
  const normalizedEmailAddress = credentials.emailAddress.trim()

  if (!normalizedEmailAddress) {
    nextFieldErrors.emailAddress = 'Email wajib diisi.'
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(normalizedEmailAddress)) {
    nextFieldErrors.emailAddress = 'Format email belum valid.'
  }

  if (!credentials.password) {
    nextFieldErrors.password = 'Kata sandi wajib diisi.'
  } else if (credentials.password.length < 8) {
    nextFieldErrors.password = 'Kata sandi minimal 8 karakter.'
  }

  return nextFieldErrors
}

function createAuthenticatedCatalogSession(
  credentials: CatalogAccountCredentials,
  loginApiResponse: LoginApiResponse | null,
): AuthenticatedCatalogSession {
  return {
    accountEmail: loginApiResponse?.user?.email ?? credentials.emailAddress.trim(),
    accessToken: loginApiResponse?.access_token ?? loginApiResponse?.token,
    displayName: loginApiResponse?.user?.name,
  }
}

async function parseLoginApiResponse(response: Response) {
  const responseText = await response.text()

  if (!responseText) {
    return null
  }

  return JSON.parse(responseText) as LoginApiResponse
}

export function useCatalogAccountLogin({
  onAuthenticatedCatalogSession,
}: UseCatalogAccountLoginOptions = {}) {
  const [rememberedAccountEmail] = useState(() => readRememberedCatalogAccount())
  const [catalogAccountCredentials, setCatalogAccountCredentials] =
    useState<CatalogAccountCredentials>({
      emailAddress: rememberedAccountEmail,
      password: '',
      shouldRememberAccount: rememberedAccountEmail.length > 0,
    })
  const [credentialFieldErrors, setCredentialFieldErrors] =
    useState<LoginFieldErrors>({})
  const [loginErrorMessage, setLoginErrorMessage] = useState<string | null>(null)
  const [isSubmittingLogin, setIsSubmittingLogin] = useState(false)

  const hasRememberedCatalogAccount = rememberedAccountEmail.length > 0
  const isLoginActionDisabled =
    isSubmittingLogin ||
    !catalogAccountCredentials.emailAddress.trim() ||
    !catalogAccountCredentials.password

  const updateCatalogLoginEmail = useCallback((emailAddress: string) => {
    setCatalogAccountCredentials((currentCredentials) => ({
      ...currentCredentials,
      emailAddress,
    }))
    setCredentialFieldErrors((currentErrors) => ({
      ...currentErrors,
      emailAddress: undefined,
    }))
    setLoginErrorMessage(null)
  }, [])

  const updateCatalogLoginPassword = useCallback((password: string) => {
    setCatalogAccountCredentials((currentCredentials) => ({
      ...currentCredentials,
      password,
    }))
    setCredentialFieldErrors((currentErrors) => ({
      ...currentErrors,
      password: undefined,
    }))
    setLoginErrorMessage(null)
  }, [])

  const updateRememberCatalogAccountPreference = useCallback(
    (shouldRememberAccount: boolean) => {
      setCatalogAccountCredentials((currentCredentials) => ({
        ...currentCredentials,
        shouldRememberAccount,
      }))
    },
    [],
  )

  const submitCatalogAccountLogin = useCallback(async () => {
    const nextFieldErrors = validateCatalogAccountCredentials(catalogAccountCredentials)

    if (Object.keys(nextFieldErrors).length > 0) {
      setCredentialFieldErrors(nextFieldErrors)
      setLoginErrorMessage(null)
      return
    }

    setIsSubmittingLogin(true)
    setLoginErrorMessage(null)

    try {
      const loginResponse = await fetch(ACCOUNT_LOGIN_ENDPOINT, {
        body: JSON.stringify({
          email: catalogAccountCredentials.emailAddress.trim(),
          password: catalogAccountCredentials.password,
        }),
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        method: 'POST',
      })

      if (!loginResponse.ok) {
        throw new Error(`Login API responded with ${loginResponse.status}`)
      }

      const loginApiResponse = await parseLoginApiResponse(loginResponse)
      const authenticatedCatalogSession = createAuthenticatedCatalogSession(
        catalogAccountCredentials,
        loginApiResponse,
      )

      persistRememberedCatalogAccount(catalogAccountCredentials)
      persistAuthenticatedCatalogSession(authenticatedCatalogSession)
      onAuthenticatedCatalogSession?.(authenticatedCatalogSession)
    } catch {
      setLoginErrorMessage(friendlyLoginErrorMessage)
    } finally {
      setIsSubmittingLogin(false)
    }
  }, [catalogAccountCredentials, onAuthenticatedCatalogSession])

  return {
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
  }
}
