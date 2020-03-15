import React from "react";
import ReactDOM from "react-dom";
import axios from "axios";
import { BrowserRouter as Router, Route } from "react-router-dom";
import Cookies from "js-cookie";
import App from "./js/pages/App";
import Top from "./js/pages/Top";
import Signup from "./js/pages/Signup";
import Signin from "./js/pages/Signin";
import Timeline from "./js/pages/Timeline";
import Mypage from "./js/pages/Mypage";
import CreatePost from "./js/pages/CreatePost";
import * as serviceWorker from "./serviceWorker";

class AppFrame extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isSignedIn: Cookies.get("session_id") !== ""
    };
    this.handleAuth = this.handleAuth.bind(this);
    this.signOut = this.signOut.bind(this);
  }

  handleAuth() {
    this.setState({
      isSignedIn: Cookies.get("session_id") !== ""
    });
  }

  signOut() {
    axios.defaults.withCredentials = true;
    axios.get("http://localhost:8088/users/signout").then(() => {
      this.handleAuth();
    });
  }

  render() {
    return (
      <App isSignedIn={this.state.isSignedIn} signOut={this.signOut}>
        <Route exact path="/" component={Top}></Route>
        <Route
          path="/signup"
          render={() => <Signup signIn={this.handleAuth} />}
        ></Route>
        <Route
          path="/signin"
          render={() => <Signin signIn={this.handleAuth} />}
        ></Route>
        <Route path="/timeline" component={Timeline}></Route>
        <Route path="/mypage" component={Mypage}></Route>
        <Route path="/create" component={CreatePost}></Route>
      </App>
    );
  }
}

ReactDOM.render(
  <Router>
    <AppFrame />
  </Router>,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
