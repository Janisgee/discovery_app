import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faFileLines,
  faHeart,
  faMagnifyingGlass,
} from '@fortawesome/free-solid-svg-icons';

import { Button } from '@/app/ui/buttons';

export default async function Home() {
  return (
    <div className='font-inter background-yellow'>
      <div className='py-5 px-10'>
        <div className='block-center'>
          <div className='w-full max-w-sm flex justify-between'>
            <FontAwesomeIcon icon={faFileLines} size='3x' />
            <FontAwesomeIcon icon={faHeart} size='3x' />
          </div>
        </div>
      </div>
      <div className='bg-white rounded-t-3xl '>
        <div className='px-10 pb-10'>
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
            <form className=''>
              <div className='block-center flex-row'>
                <FontAwesomeIcon icon={faMagnifyingGlass} size='2x' />
                <label
                  htmlFor='default-search'
                  className='mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white'
                >
                  Search
                </label>
                <input
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
                <div class='absolute inset-0 bg-gray-700 opacity-40 rounded-lg'></div>
                <div class='absolute inset-0 flex items-center justify-center'>
                  <h2 class='text-white'>Japan</h2>
                </div>
              </div>
              <div className='relative max-w-xl mx-auto'>
                <img
                  src='/place_img/paris-france.jpg'
                  className='h-32 w-full object-cover mt-5 rounded-lg'
                  alt='Image of place'
                />
                <div class='absolute inset-0 bg-gray-700 opacity-40 rounded-lg'></div>
                <div class='absolute inset-0 flex items-center justify-center'>
                  <h2 class='text-white'>Korea</h2>
                </div>
              </div>
              <div className='relative max-w-xl mx-auto'>
                <img
                  src='/place_img/paris-france.jpg'
                  className='h-32 w-full object-cover mt-5 rounded-lg'
                  alt='Image of place'
                />
                <div class='absolute inset-0 bg-gray-700 opacity-40 rounded-lg'></div>
                <div class='absolute inset-0 flex items-center justify-center'>
                  <h2 class='text-white'>France</h2>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
