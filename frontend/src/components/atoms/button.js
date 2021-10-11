import React from "react";

export class Button extends React.Component {
  render() {
    return (
      <button onClick={this.props.onClick}>
        {this.props.label}
      </button>
    );
  }
}

Button.defaultProps = {
  onClick: () => { },
  label: "[label]",
};
