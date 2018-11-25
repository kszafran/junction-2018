import React from "react";
import logo from "../../assets/images/logo.png";
import "./index.css";

const Loader = () => {
  return (
    <div className="loader">
      <img src={logo} alt="loader logo" />
      <div className="lds-ellipsis">
        <div />
        <div />
        <div />
        <div />
      </div>
    </div>
  );
};

export default Loader;
