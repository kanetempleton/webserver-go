import ReactDOM from "react-dom";
import React, { Component } from "react";

export default class Header extends Component {
    render() {
        return (
            <div>
                <div className="title">Welcome to WebGo!</div>
                <hr className="divider" />
            </div>
        );
    }
}

ReactDOM.render(<Header />, document.getElementById("root"));
