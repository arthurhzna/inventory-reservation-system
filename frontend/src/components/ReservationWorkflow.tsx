import type { ConfirmResponse, Reservation } from '../types/inventory'
import { formatCountdown } from '../utils/time'

type ReservationWorkflowProps = {
  activeReservation: Reservation | null
  confirmedReservation: ConfirmResponse | null
  isConfirming: boolean
  remainingMs: number
  isExpired: boolean
  onConfirm: () => void
}

export function ReservationWorkflow({
  activeReservation,
  confirmedReservation,
  isConfirming,
  remainingMs,
  isExpired,
  onConfirm,
}: ReservationWorkflowProps) {
  return (
    <div className="panel workflow-panel">
      <div className="panel-header">
        <div>
          <p className="eyebrow">Workflow</p>
          <h2>Expiration and confirmation</h2>
        </div>
      </div>

      {activeReservation ? (
        <div className="reservation-card">
          <div className="reservation-id">
            <span>Reservation ID</span>
            <strong>{activeReservation.reservation_id}</strong>
          </div>

          <div className="countdown-block">
            <span>{isExpired ? 'Expired' : 'Time left'}</span>
            <strong>{formatCountdown(remainingMs)}</strong>
          </div>

          <dl className="reservation-details">
            <div>
              <dt>Item</dt>
              <dd>{activeReservation.item_id}</dd>
            </div>
            <div>
              <dt>Quantity</dt>
              <dd>{activeReservation.quantity}</dd>
            </div>
            <div>
              <dt>Expires</dt>
              <dd>{new Date(activeReservation.expires_at).toLocaleTimeString()}</dd>
            </div>
          </dl>

          {isExpired ? (
            <p className="inline-error">
              Reservation expired. Request a new reservation before confirming
              purchase.
            </p>
          ) : null}

          <button
            className="primary-button"
            type="button"
            onClick={onConfirm}
            disabled={isConfirming || isExpired}
          >
            {isConfirming ? 'Confirming' : 'Confirm Purchase'}
          </button>
        </div>
      ) : (
        <div className="empty-state">
          <strong>No active reservation</strong>
          <span>Submit a reservation to start the countdown.</span>
        </div>
      )}

      {confirmedReservation ? (
        <div className="confirmation">
          <span>Confirmed</span>
          <strong>{confirmedReservation.reservation_id}</strong>
          <small>
            {new Date(confirmedReservation.confirmed_at).toLocaleString()}
          </small>
        </div>
      ) : null}
    </div>
  )
}
