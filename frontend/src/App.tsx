import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'
import { CatalogPage } from './features/catalog/CatalogPage'
import { LoginPage } from './features/login/LoginPage'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navigate to="/login" replace />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/katalog" element={<CatalogPage />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
