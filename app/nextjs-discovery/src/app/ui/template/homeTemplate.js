import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import { faFileLines, faHeart } from '@fortawesome/free-solid-svg-icons';

export default function AppTemplate({ children }) {
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
        <div className='px-10 pb-10'>{children}</div>
      </div>
    </div>
  );
}
