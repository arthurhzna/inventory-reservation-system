import { apiClient } from './apiClient'
import type {
  ConfirmResponse,
  InventoryListResponse,
  Reservation,
  ReserveInventoryPayload,
} from '../types/inventory'

export function getInventory() {
  return apiClient<InventoryListResponse>('/api/v1/inventory')
}

export function reserveInventory(payload: ReserveInventoryPayload) {
  return apiClient<Reservation>('/api/v1/inventory/reserve', {
    method: 'POST',
    body: JSON.stringify(payload),
  })
}

export function confirmReservation(reservationId: string) {
  return apiClient<ConfirmResponse>('/api/v1/inventory/confirm', {
    method: 'POST',
    body: JSON.stringify({
      reservation_id: reservationId,
    }),
  })
}
