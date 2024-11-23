import Link from 'next/link';

export function Button({ useFor, link, color }) {
  return (
    <Link href={link}>
      <button className={`font-space_mono ${color}`}>{useFor}</button>
    </Link>
  );
}
