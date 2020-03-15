import React from "react";
import axios from "axios";
import { withRouter } from "react-router";
import Form from "../components/Form";

class Signup extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      formContents: [
        { label: "User ID", name: "user_id", type: "text" },
        { label: "PassWord", name: "password", type: "text" },
        { label: "Name", name: "name", type: "text" }
      ]
    };
    this.handleFormSubmit = this.handleFormSubmit.bind(this);
  }

  handleFormSubmit(event) {
    event.preventDefault();
    var params = {};
    for (var i = 0; i < this.state.formContents.length; i++) {
      const formName = this.state.formContents[i].name;
      params[formName] = this.elements[formName].value;
    }
    axios.defaults.withCredentials = true;
    axios.post("http://localhost:8088/users/signup", params).then(() => {
      this.props.history.push("/mypage");
      this.props.signIn();
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
export default withRouter(Signup);
