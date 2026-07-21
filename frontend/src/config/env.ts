declare const __URL_API_INVENTORY__: string

function normalizeApiUrl(value: string) {
  const trimmedValue = value.trim()

  if (!trimmedValue) {
    throw new Error('URL_API_INVENTORY is not configured')
  }

  return trimmedValue.replace(/\/$/, '')
}

export const API_BASE_URL = normalizeApiUrl(__URL_API_INVENTORY__)