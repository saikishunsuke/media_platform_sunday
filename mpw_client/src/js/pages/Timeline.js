import React from "react";
import axios from "axios";
import Card from "../components/Card";
import "../../css/Timeline.scss";

export default class Timeline extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      contents: []
    };
    this.getPosts();
  }

  getPosts() {
    axios.get("http://localhost:8088/post/index").then(response => {
      this.setState({
        contents: response.data ? response.data.reverse() : []
      });
    });
  }

  render() {
    const cards = this.state.contents.map((content, index) => {
      return <Card content={content} key={index} />;
    });
    return <div className="cards">{cards}</div>;
  }
}
