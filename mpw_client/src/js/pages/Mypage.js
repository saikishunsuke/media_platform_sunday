import React from "react";
import axios from "axios";
import Card from "../components/Card";
import "../../css/Mypage.scss";

axios.defaults.withCredentials = true;

export default class Mypage extends React.Component {
  constructor(props) {
    super(props);
    this.state = { user: Object, contents: [] };
    this.getPosts();
    this.getUser();
  }

  getPosts() {
    axios
      .get("http://localhost:8088/post/mine", { withCredentials: true })
      .then(response => {
        this.setState({
          contents: response.data ? response.data.reverse() : []
        });
      });
  }

  getUser() {
    axios.get("http://localhost:8088/users/sign_in_user").then(response => {
      console.log(response.data.Name);
      this.setState({
        user: response.data
      });
    });
  }

  render() {
    const cards = this.state.contents.map((content, index) => {
      return <Card content={content} key={index} />;
    });
    return (
      <div className="mypage">
        <p className="user title is-5">
          {this.state.user.Name}
          <span className="user-id">@{this.state.user.user_id}</span>
        </p>
        <div className="cards">{cards}</div>
      </div>
    );
  }
}
