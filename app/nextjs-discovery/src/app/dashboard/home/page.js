'use client';
import { useRouter } from 'next/navigation';
// import { useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons';

import { Button } from '@/app/ui/buttons';
import HomeTemplate from '@/app/ui/template/homeTemplate';

export default function Home() {
  const router = useRouter();

  const handleSearchSubmit = (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const searchData = formData.get('search');
    if (!searchData) {
      alert('Please enter a location');
      return;
    }
    router.push(`/dashboard/location/${encodeURIComponent(searchData)}`);
  };

  return (
    <div>
      <HomeTemplate>
        <div className='block-center flex-col pb-10'>
          <img
            src='/user_img/default.jpg'
            className='mb-5 h-32 rounded-full'
            alt='default profile picture'
          />
          <h6 className=''>Janis Chan</h6>
          <p className='text-color-dark_grey'>7 bookmark places</p>
        </div>
        <div className='block-center flex-col'>
          <Button useFor='✒️ Plan New Trip' link='/' color='btn-violet' />

          <form onSubmit={handleSearchSubmit}>
            <div className='block-center flex-row'>
              <FontAwesomeIcon icon={faMagnifyingGlass} size='2x' />
              <label
                htmlFor='default-search'
                className='mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white'
              >
                Search
              </label>
              <input
                name='search'
                type='search'
                id='default-search'
                className='block w-full ml-2 mt-2 p-2 ps-5 text-sm text-gray-900 border border-gray-300 rounded-full bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
                placeholder='Type location to search'
                required
              />
            </div>
          </form>
        </div>
        <div className='my-5 block-center flex-col '>
          <h3>Popular Destination</h3>
          {/* <div className='w-full h-96 overflow-hidden inline-block rounded-lg'> */}
          <div className='w-full h-96 overflow-auto rounded-lg'>
            <div className='relative max-w-xl mx-auto'>
              <img
                src='/place_img/paris-france.jpg'
                className='h-32 w-full object-cover mt-5 rounded-lg'
                alt='Image of place'
              />
              <div className='absolute inset-0 bg-gray-700 opacity-40 rounded-lg'></div>
              <div className='absolute inset-0 flex items-center justify-center'>
                <h2 className='text-white'>Japan</h2>
              </div>
            </div>
            <div className='relative max-w-xl mx-auto'>
              <img
                src='/place_img/paris-france.jpg'
                className='h-32 w-full object-cover mt-5 rounded-lg'
                alt='Image of place'
              />
              <div className='absolute inset-0 bg-gray-700 opacity-40 rounded-lg'></div>
              <div className='absolute inset-0 flex items-center justify-center'>
                <h2 className='text-white'>Korea</h2>
              </div>
            </div>
            <div className='relative max-w-xl mx-auto'>
              <img
                src='/place_img/paris-france.jpg'
                className='h-32 w-full object-cover mt-5 rounded-lg'
                alt='Image of place'
              />
              <div className='absolute inset-0 bg-gray-700 opacity-40 rounded-lg'></div>
              <div className='absolute inset-0 flex items-center justify-center'>
                <h2 className='text-white'>France</h2>
              </div>
            </div>
          </div>
        </div>
      </HomeTemplate>
    </div>
  );
}
