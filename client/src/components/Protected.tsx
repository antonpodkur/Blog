import { ReactNode } from "react"
import { Navigate } from "react-router-dom"

interface ProtectedPros {
    isLoggedIn: boolean
    children: ReactNode
}
const Protected: React.FC<ProtectedPros> = ({isLoggedIn, children}) => {
    if (!isLoggedIn) {
        return <Navigate to="/login" replace />
    }
    return children
} 

export default Protected