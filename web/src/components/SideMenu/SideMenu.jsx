import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { LinksPathnames } from '../../constant/commonConstant';
import './style.scss';

export const SideMenu = () => {
  const [isOpen, setIsOpen] = useState(false);

  const pathname = useSelector(({ router }) => router.location.pathname);

  useEffect(() => setIsOpen(false), [pathname]);

  const links = [
    {
      label: LinksPathnames.HOME.name,
      url: LinksPathnames.HOME.path,
    },
    {
      label: LinksPathnames.IMAGE.name,
      url: LinksPathnames.IMAGE.path,
    },
    {
      label: LinksPathnames.TEST.name,
      url: LinksPathnames.TEST.path,
    },
  ];

  const linksList = links => {
    const linkArray = links.map(item => (
      <li key={item.url}>
        <Link to={item.url}>{item.label}</Link>
      </li>
    ));
    return linkArray;
  };

  return (
    <aside>
      <nav>
        <input
          type="checkbox"
          className="check"
          id="checked"
          checked={isOpen}
          onChange={() => setIsOpen(!isOpen)}
        />
        <label className="menu-btn" htmlFor="checked">
          <span
            className={
              isOpen
                ? 'hamburger hamburger--3dx is-active'
                : 'hamburger hamburger--3dx'
            }
          >
            <span className="hamburger-box">
              <span className="hamburger-inner" />
            </span>
          </span>
        </label>
        <label className="close-menu" htmlFor="checked" />
        <div className="drawer-menu">
          <ul className="sub">{linksList(links)}</ul>
        </div>
      </nav>
    </aside>
  );
};
