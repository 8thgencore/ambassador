import "./App.css";
import ProductsFrontend from "./pages/ProductsFrontend";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Profile from "./pages/Profile";
import Stats from "./pages/Stats";
import Rankings from "./pages/Rankings";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path={"/"} element={<ProductsFrontend />} />
        <Route path={"/login"} element={<Login />} />
        <Route path={"/register"} element={<Register />} />
        <Route path={"/profile"} element={<Profile />} />
        <Route path={"/stats"} element={<Stats />} />
        <Route path={"/rankings"} element={<Rankings />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
