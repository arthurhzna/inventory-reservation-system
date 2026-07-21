import type { FormEvent } from 'react'
import type { InventoryItem } from '../types/inventory'

type ReservationControlProps = {
  userId: string
  quantity: number
  selectedItem: InventoryItem | null
  isReserving: boolean
  onUserIdChange: (userId: string) => void
  onQuantityChange: (quantity: number) => void
  onSubmit: (event: FormEvent<HTMLFormElement>) => void
}

export function ReservationControl({
  userId,
  quantity,
  selectedItem,
  isReserving,
  onUserIdChange,
  onQuantityChange,
  onSubmit,
}: ReservationControlProps) {
  return (
    <div className="panel">
      <div className="panel-header">
        <div>
          <p className="eyebrow">Reservation Control</p>
          <h2>Reserve stock</h2>
        </div>
      </div>

      <form className="reservation-form" onSubmit={onSubmit}>
        <label className="field">
          User ID
          <input
            value={userId}
            onChange={(event) => onUserIdChange(event.target.value)}
            placeholder="user_001"
            required
          />
        </label>

        <label className="field">
          Quantity
          <input
            min="1"
            type="number"
            value={quantity}
            onChange={(event) => onQuantityChange(Number(event.target.value))}
            required
          />
        </label>

        <button
          className="primary-button"
          type="submit"
          disabled={isReserving || !selectedItem}
        >
          {isReserving ? 'Reserving' : 'Request Reservation'}
        </button>
      </form>
    </div>
  )
}
