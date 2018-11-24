import React from "react";
import avatarPlaceholder from "../../assets/images/avatar.png";

import "./index.css";

const Header = () => {
  return (
    <div className="header">
      <div />
      <h1>Cisco comfort</h1>
      <img src={avatarPlaceholder} alt="" />
    </div>
  );
};

export default Header;
