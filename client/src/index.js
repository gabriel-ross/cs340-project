import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Pokemon from './App';
import Generations from './pages/Generations';
import Moves from './pages/Moves';
import Types from './pages/Types';
import PokemonMoves from './pages/PokemonMoves';
import reportWebVitals from './reportWebVitals';
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import {
  BrowserRouter,
  Routes,
  Route
} from "react-router-dom";

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
    <Routes>
      <Route path="/" element={<Pokemon />} />
      <Route path="moves" element={<Moves />} />
      <Route path="types" element={<Types />} />
      <Route path="generations" element={<Generations />} />
      <Route path="pokemonmoves" element={<PokemonMoves />} />
    </Routes>
  </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
