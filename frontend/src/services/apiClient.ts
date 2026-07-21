import { API_BASE_URL } from '../config/env'

type ApiResponse<T> = {
  message?: string
  data?: T
  errors?: Array<{
    field?: string
    message?: string
  }>
}

function getErrorMessage<T>(payload: ApiResponse<T>, status: number) {
  const validationMessage = payload.errors
    ?.map((error) =>
      error.field && error.message
        ? `${error.field}: ${error.message}`
        : error.message,
    )
    .filter(Boolean)
    .join(', ')

  return (
    validationMessage ||
    payload.message ||
    `Request failed with status ${status}`
  )
}

export async function apiClient<T>(
  path: string,
  options?: RequestInit,
): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${path}`, {
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
    ...options,
  })

  const payload = (await response.json().catch(() => ({}))) as ApiResponse<T>

  if (!response.ok) {
    throw new Error(getErrorMessage(payload, response.status))
  }

  if (payload.data === undefined) {
    throw new Error('Response data is empty')
  }

  return payload.data
}
