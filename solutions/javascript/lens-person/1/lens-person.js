import { Person } from './person';
import { Name } from './name';
import { Born } from './born';
import { Address } from './address';
import { Lens } from './lens';

export const nameLens = new Lens(
  person => person.name,
  (person, name) => ({...person, name}),
)

export const bornAtLens = new Lens(
  person => person.born.bornAt,
  (person, bornAt) => ({...person, born: {...person.born, bornAt}}),
)

export const streetLens = new Lens(
  person => person.address.street,
  (person, street) => ({...person, address: {...person.address, street}}),
)