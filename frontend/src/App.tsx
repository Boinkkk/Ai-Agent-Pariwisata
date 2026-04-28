import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'
import { CatalogPage } from './features/catalog/CatalogPage'
import { ProductDetailPage } from './features/catalog/ProductDetailPage'
import { LoginPage } from './features/login/LoginPage'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navigate to="/login" replace />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/katalog" element={<CatalogPage />} />
        <Route path="/katalog/:productSlug" element={<ProductDetailPage />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
