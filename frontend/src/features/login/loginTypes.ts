export interface CatalogAccountCredentials {
  emailAddress: string
  password: string
  shouldRememberAccount: boolean
}

export interface LoginFieldErrors {
  emailAddress?: string
  password?: string
}

export interface AuthenticatedCatalogSession {
  accountEmail: string
  accessToken?: string
  displayName?: string
}

export interface LoginApiResponse {
  access_token?: string
  token?: string
  user?: {
    email?: string
    name?: string
  }
}
