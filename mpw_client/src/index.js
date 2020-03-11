import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router, Route } from "react-router-dom";
import App from "./js/pages/App";
import Timeline from "./js/pages/Timeline";
import Mypage from "./js/pages/Mypage";
import * as serviceWorker from "./serviceWorker";

// ReactDOM.render(<App />, document.getElementById("root"));
ReactDOM.render(
  <Router>
    <App>
      {/* <Route exact path="/" component={Top}></Route> */}
      <Route path="/timeline" component={Timeline}></Route>
      <Route path="/mypage" component={Mypage}></Route>
    </App>
  </Router>,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
