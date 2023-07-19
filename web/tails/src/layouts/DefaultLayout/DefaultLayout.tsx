import React, { ReactNode } from 'react';
import logo from './logo.svg';

import { Route, Routes } from 'react-router-dom';

import "./DefaultLayout.scss"

interface IDefaultLayoutProps {
    children: ReactNode
}

export default (props: IDefaultLayoutProps) => {
    return (
        <div className="default-layout">
            <nav>asdf</nav>
            {/* <div className="content">
                { props.children }
            </div> */}
        </div>
    )
}
