import React from "react";
import avatarPlaceholder from "../../assets/images/avatar.png";
import { Link } from "react-router-dom";

import "./index.css";

const Header = () => {
  return (
    <div className="header">
      <div />
      <Link to="/">
        <h1>Cisco Comfort</h1>
      </Link>
      <img src={avatarPlaceholder} alt="" />
    </div>
  );
};

export default Header;
