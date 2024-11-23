import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import {
  faFileLines,
  faHeart,
  faHouse,
} from '@fortawesome/free-solid-svg-icons';

import { Button } from '@/app/ui/buttons';

export default function AppTemplate({ children }) {
  return (
    <div className='font-inter background-yellow'>
      <div className='py-5 px-5'>
        <div className='block-center'>
          <div className='w-full max-w-sm flex justify-between'>
            <div>
              <span className='mr-5'>
                <FontAwesomeIcon icon={faHouse} size='3x' />
              </span>
              <FontAwesomeIcon icon={faFileLines} size='3x' />
            </div>
            <div className='flex items-center justify-center'>
              <FontAwesomeIcon icon={faHeart} size='3x' />
              <span className=' ml-2 inline items-start text-xl'>
                <Button useFor='JG' link='/dashboard/home' color='btn-grey' />
              </span>
            </div>
          </div>
        </div>
      </div>
      <div className='bg-white rounded-t-3xl '>
        <div className='px-10 pb-10'>{children}</div>
      </div>
    </div>
  );
}
