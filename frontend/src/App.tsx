import { useCallback, useEffect, useMemo, useState } from 'react'
import type { FormEvent } from 'react'
import './App.css'
import { DashboardHeader } from './components/DashboardHeader'
import { InventoryTracker } from './components/InventoryTracker'
import { NoticeBanner } from './components/NoticeBanner'
import { ReservationControl } from './components/ReservationControl'
import { ReservationWorkflow } from './components/ReservationWorkflow'
import { API_BASE_URL } from './config/env'
import {
  confirmReservation,
  getInventory,
  reserveInventory,
} from './services/inventoryService'
import type {
  ConfirmResponse,
  InventoryItem,
  Notice,
  Reservation,
} from './types/inventory'

const POLL_INTERVAL_MS = 5000
const DEFAULT_ITEM_ID = 'item_001'

function App() {
  const [inventory, setInventory] = useState<InventoryItem[]>([])
  const [selectedItemId, setSelectedItemId] = useState(DEFAULT_ITEM_ID)
  const [userId, setUserId] = useState('user_001')
  const [quantity, setQuantity] = useState(1)
  const [activeReservation, setActiveReservation] = useState<Reservation | null>(
    null,
  )
  const [confirmedReservation, setConfirmedReservation] =
    useState<ConfirmResponse | null>(null)
  const [notice, setNotice] = useState<Notice | null>(null)
  const [isLoadingInventory, setIsLoadingInventory] = useState(false)
  const [isReserving, setIsReserving] = useState(false)
  const [isConfirming, setIsConfirming] = useState(false)
  const [lastUpdatedAt, setLastUpdatedAt] = useState<Date | null>(null)
  const [now, setNow] = useState(() => Date.now())

  const selectedItem = useMemo(
    () => inventory.find((item) => item.item_id === selectedItemId) ?? null,
    [inventory, selectedItemId],
  )

  const reservationRemainingMs = activeReservation
    ? new Date(activeReservation.expires_at).getTime() - now
    : 0
  const isReservationExpired = Boolean(
    activeReservation && reservationRemainingMs <= 0,
  )

  const loadInventory = useCallback(
    async (mode: 'silent' | 'manual' = 'silent') => {
      if (mode === 'manual') {
        setNotice(null)
      }

      setIsLoadingInventory(true)

      try {
        const data = await getInventory()
        const items = data.items ?? []

        setInventory(items)
        setLastUpdatedAt(new Date())
        setSelectedItemId((currentItemId) =>
          items.some((item) => item.item_id === currentItemId)
            ? currentItemId
            : (items[0]?.item_id ?? DEFAULT_ITEM_ID),
        )

        if (mode === 'manual') {
          setNotice({ type: 'success', message: 'Inventory refreshed.' })
        }
      } catch (error) {
        setNotice({
          type: 'error',
          message:
            error instanceof Error
              ? error.message
              : 'Unable to load inventory.',
        })
      } finally {
        setIsLoadingInventory(false)
      }
    },
    [],
  )

  useEffect(() => {
    const timeoutId = window.setTimeout(() => {
      void loadInventory()
    }, 0)
    const intervalId = window.setInterval(() => {
      void loadInventory()
    }, POLL_INTERVAL_MS)

    return () => {
      window.clearTimeout(timeoutId)
      window.clearInterval(intervalId)
    }
  }, [loadInventory])

  useEffect(() => {
    const intervalId = window.setInterval(() => setNow(Date.now()), 1000)
    return () => window.clearInterval(intervalId)
  }, [])

  async function handleReserve(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setNotice(null)
    setConfirmedReservation(null)
    setIsReserving(true)

    try {
      const reservation = await reserveInventory({
        user_id: userId,
        item_id: selectedItemId,
        quantity,
      })

      setActiveReservation(reservation)
      setNotice({
        type: 'success',
        message: `Reserved ${reservation.quantity} unit(s) for ${reservation.item_id}.`,
      })
      await loadInventory()
    } catch (error) {
      setNotice({
        type: 'error',
        message:
          error instanceof Error ? error.message : 'Reservation request failed.',
      })
    } finally {
      setIsReserving(false)
    }
  }

  async function handleConfirm() {
    if (!activeReservation) return

    setNotice(null)
    setIsConfirming(true)

    try {
      const confirmation = await confirmReservation(
        activeReservation.reservation_id,
      )

      setConfirmedReservation(confirmation)
      setActiveReservation(null)
      setNotice({ type: 'success', message: 'Purchase confirmed.' })
      await loadInventory()
    } catch (error) {
      setNotice({
        type: 'error',
        message:
          error instanceof Error
            ? error.message
            : 'Unable to confirm reservation.',
      })
      await loadInventory()
    } finally {
      setIsConfirming(false)
    }
  }

  return (
    <main className="dashboard-shell">
      <DashboardHeader apiBaseUrl={API_BASE_URL} />
      <NoticeBanner notice={notice} />

      <section className="dashboard-grid">
        <InventoryTracker
          inventory={inventory}
          selectedItem={selectedItem}
          selectedItemId={selectedItemId}
          isLoading={isLoadingInventory}
          lastUpdatedAt={lastUpdatedAt}
          onRefresh={() => void loadInventory('manual')}
          onSelectItem={setSelectedItemId}
        />

        <ReservationControl
          userId={userId}
          quantity={quantity}
          selectedItem={selectedItem}
          isReserving={isReserving}
          onUserIdChange={setUserId}
          onQuantityChange={setQuantity}
          onSubmit={handleReserve}
        />

        <ReservationWorkflow
          activeReservation={activeReservation}
          confirmedReservation={confirmedReservation}
          isConfirming={isConfirming}
          remainingMs={reservationRemainingMs}
          isExpired={isReservationExpired}
          onConfirm={() => void handleConfirm()}
        />
      </section>
    </main>
  )
}

export default App
