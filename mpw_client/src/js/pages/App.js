import React from "react";
import Sidebar from "../components/Sidebar";
import "../../css/App.scss";

export default class App extends React.Component {
  render() {
    return (
      <div className="columns">
        <Sidebar
          isSignedIn={this.props.isSignedIn}
          signOut={() => this.props.signOut()}
        />
        <div className="column main-content">{this.props.children}</div>
      </div>
    );
  }
}
