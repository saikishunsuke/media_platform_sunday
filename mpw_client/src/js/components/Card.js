import React from "react";
import "../../css/Card.scss";

export default class Card extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isActive: false
    };
    this.toggleActive = this.toggleActive.bind(this);
  }

  toggleActive() {
    const newState = { isActive: !this.state.isActive };
    this.setState(newState);
  }

  render() {
    return (
      <div
        className={`card ${this.state.isActive ? "active" : "non-active"}`}
        onClick={this.toggleActive}
      >
        <header className="card-header">
          <div className="media">
            <div className="media-left">
              <figure className="image">
                <img
                  src="https://bulma.io/images/placeholders/96x96.png"
                  alt="Placeholder image"
                />
              </figure>
            </div>
            <div className="media-content">
              <p className="user-name">
                {this.props.content.User.Name}
                <span className="user-id">
                  @{this.props.content.User.user_id}
                </span>
              </p>
            </div>
          </div>
        </header>
        <div className="card-content">
          <div className="content">
            <p className="title is-5">{this.props.content.Title}</p>
            <p className="text">{this.props.content.Text}</p>
            <a href="#">#css</a>
            <a href="#">#responsive</a>
            <br />
            <time dateTime="2016-1-1">{this.props.content.CreatedAt}</time>
          </div>
        </div>
      </div>
    );
  }
}
