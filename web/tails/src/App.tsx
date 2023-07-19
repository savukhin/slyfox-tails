import React from 'react';
import logo from './logo.svg';

import { Route, Routes } from 'react-router-dom';

import './App.scss';
import DefaultLayout from './layouts/DefaultLayout/DefaultLayout';

function App() {
  return (
    <div className="App">
      <DefaultLayout>
          <Routes>
              {/* <Route path="/" element={<IndexPage/>} />
              <Route path="/contacts" element={<ContactPage/>} />
              <Route path="/about" element={<AboutUsPage/>} />
              <Route path="/order" element={<OrderPage/>} />
              <Route path="/capabilities/:title" element={<CapabilityPage/>} />
              <Route path="/create-post" element={<CreatePostPage/>} />
              <Route path="/blog" element={<BlogPage/>} />
              <Route path="/blog/page/:page" element={<BlogPage/>} />
              
              <Route path="/map" element={<RedirectPage location="https://yandex.ru/maps/-/CCUzqJwmsC"/> } />

              <Route path="*" element={<NotFoundPage />} /> */}
          </Routes>
      </DefaultLayout>
    </div>
  );
}

export default App;
