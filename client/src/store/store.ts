import {create} from 'zustand'
import {persist} from 'zustand/middleware'
import { User } from '../models/user'

type State = {
    isLoggedIn: boolean,
    user: User | null,
}

type Actions = {
    setLoggedIn: (loggedIn: boolean) => void,
    setUser: (user: User) => void,
    reset: () => void
}

const initialState: State = {
    isLoggedIn: false,
    user: null,
}

export const useAuthStore = create<State & Actions>()(
    persist(
        (set) => ({
            ...initialState,
            setLoggedIn: (loggedIn: boolean) => set(() => ({isLoggedIn: loggedIn})),
            setUser: (user: User) => set(() => ({user: user})),
            reset: () => set(initialState)
        }),
        {
            name: "auth-storage"
        }
    )
)
