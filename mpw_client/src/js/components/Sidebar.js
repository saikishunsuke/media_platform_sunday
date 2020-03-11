import React from "react";
import { Link } from "react-router-dom";
import "../../css/SideBar.scss";

export default class Sidebar extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      signedIn: [
        {
          text: "Top",
          link: ""
        },
        {
          text: "Mypage",
          link: "mypage"
        },
        {
          text: "Timeline",
          link: "timeline"
        },
        {
          text: "Sign out",
          link: "",
          onClick: this.props.signOut
        }
      ],
      notSignedIn: [
        {
          text: "Top",
          link: ""
        },
        {
          text: "Sign in",
          link: "signin",
          onClick: this.props.signIn
        },
        {
          text: "Sign up",
          link: "signup",
          onClick: this.props.signIn
        }
      ]
    };
  }
  render() {
    const items = this.props.isSignedIn
      ? this.state.signedIn
      : this.state.notSignedIn;
    return (
      <div className="column is-one-fifth sidebar">
        {items.map(item => {
          return <SidebarItem item={item} />;
        })}
      </div>
    );
  }
}

function SidebarItem(props) {
  return (
    <Link to={props.item.link} onClick={props.item.onClick}>
      {props.item.text}
    </Link>
  );
}
