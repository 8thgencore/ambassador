import "./App.css";
import ProductsFrontend from "./pages/ProductsFrontend";
import { BrowserRouter, Route, Routes } from "react-router-dom";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path={"/"} element={<ProductsFrontend />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
