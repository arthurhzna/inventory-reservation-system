export type InventoryItem = {
  item_id: string
  item_name: string
  total_stock: number
  reserved_stock: number
  available_stock: number
}

export type InventoryListResponse = {
  items: InventoryItem[]
}

export type Reservation = {
  reservation_id: string
  item_id: string
  quantity: number
  expires_at: string
}

export type ConfirmResponse = {
  reservation_id: string
  confirmed_at: string
}

export type Notice = {
  type: 'success' | 'error' | 'info'
  message: string
}

export type ReserveInventoryPayload = {
  user_id: string
  item_id: string
  quantity: number
}
