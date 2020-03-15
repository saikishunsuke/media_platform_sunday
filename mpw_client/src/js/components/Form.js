import React from "react";
import "../../css/Form.scss";

export default class Form extends React.Component {
  constructor(props) {
    super(props);
    this.handleFormSubmit = props.onSubmit.bind(this);
  }

  render() {
    const form = this.props.formContents.map((content, index) => {
      return <FormContent content={content} key={index} />;
    });
    return (
      <form
        onSubmit={this.props.onSubmit}
        ref={this.props.myRef}
        className="is-grouped is-grouped-centered"
      >
        {form}
        <div className="button-block">
          <button className="button" type="submit">
            submit
          </button>
        </div>
      </form>
    );
  }
}

function FormContent(props) {
  let form;
  if (props.content.type === "textarea") {
    form = (
      <textarea
        className="textarea"
        name={props.content.name}
        rows="10"
      ></textarea>
    );
  } else {
    form = (
      <input
        className="input"
        type={props.content.type}
        name={props.content.name}
      />
    );
  }
  return (
    <div className="control">
      <span className="label">{props.content.label}</span>
      {form}
    </div>
  );
}
