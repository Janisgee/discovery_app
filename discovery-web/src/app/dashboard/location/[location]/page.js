'use client';

import AppTemplate from '@/app/ui/template/appTemplate';

import { useParams } from 'next/navigation';

export default function LocationPlace() {
  const params = useParams();
  console.log(params.location);
  return (
    <div>
      <AppTemplate>
        <h1> location Details: {params.location}</h1>
      </AppTemplate>
    </div>
  );
}
