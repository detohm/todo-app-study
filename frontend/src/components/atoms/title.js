import React from "react";

export class Title extends React.Component {
  render() {
    return <h1>{this.props.title}</h1>;
  }
}

Title.defaultProps = { title: "[title]" };
