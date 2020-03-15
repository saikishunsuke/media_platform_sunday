import React from "react";
import axios from "axios";
import Form from "../components/Form";

export default class CreatePost extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      formContents: [
        { label: "Title", name: "title", type: "text" },
        { label: "Text", name: "text", type: "textarea" }
      ]
    };
  }

  handleFormSubmit(event) {
    event.preventDefault();
    var params = {};
    for (var i = 0; i < this.state.formContents.length; i++) {
      const formName = this.state.formContents[i].name;
      params[formName] = this.elements[formName].value;
    }
    axios.defaults.withCredentials = true;
    axios.post("http://localhost:8088/post/new", params).then(() => {
      this.props.history.push("/mypage");
    });
  }

  render() {
    return (
      <Form
        formContents={this.state.formContents}
        onSubmit={event => {
          this.handleFormSubmit(event);
        }}
        myRef={el => (this.elements = el && el.elements)}
      />
    );
  }
}
