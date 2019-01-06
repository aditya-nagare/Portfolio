import React from "react";
import ReactDOM from "react-dom";

import "font-awesome/css/font-awesome.min.css";
import "bootstrap-css-only/css/bootstrap.min.css";
import "mdbreact/dist/css/mdb.css";

import Search from "./Search";
import "./styles.css";

function App() {
  return (
    <div className="App">
      <Search />
    </div>
  );
}

const rootElement = document.getElementById("root");
ReactDOM.render(<App />, rootElement);
