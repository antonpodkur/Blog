import axios from "axios"
import { useAuthStore } from "../store/store"

const axiosInstance = axios.create({
  baseURL: "http://localhost:4000",
  withCredentials: true
})


export const useAxios = () => {
  const isLoggedIn = useAuthStore((state) => state.isLoggedIn)
  const setLoggedIn = useAuthStore((state) => state.setLoggedIn)
  const reset = useAuthStore((state) => state.reset)

  axiosInstance.interceptors.response.use(
    (response) => {
      return response
    },
    async (error) => {
      const originalRequest = error.config
      if (
        error.response.status === 401 &&
        !originalRequest._retry &&
        isLoggedIn
      ) {
        originalRequest._retry = true
        try {
          const newResponse = await axiosInstance.get("/api/v1/auth/refresh")
          if (newResponse.status === 200) {
            setLoggedIn(true)
          } else {
            reset()
          }
          return axiosInstance(originalRequest)
        } catch (error) {
          reset()
          console.warn("Failed to refresh token ", error)
        }
      }
      return Promise.reject(error)
    }
  )

  return axiosInstance
}
