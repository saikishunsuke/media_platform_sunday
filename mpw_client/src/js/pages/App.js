import React from "react";
import Sidebar from "../components/Sidebar";
import "../../css/App.scss";

export default class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isSignedIn: false
    };
  }

  signIn() {
    this.setState({ isSignedIn: true });
  }

  signOut() {
    this.setState({ isSignedIn: false });
  }

  render() {
    return (
      <div className="columns">
        <Sidebar
          isSignedIn={this.state.isSignedIn}
          signIn={() => this.signIn()}
          signOut={() => this.signOut()}
        />
        <div className="column is-8">{this.props.children}</div>
      </div>
    );
  }
}
