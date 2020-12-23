/* eslint-disable @typescript-eslint/no-unused-vars */
const x: {
  user: {
    name: string;
    address?: {
      street: string;
      city: string;
    };
  };
} = undefined as any;

console.log(x.user.address?.city);

class Foo {
  #name: string; // name is private. diff in private and # is breakpoints wont work in #.

  constructor(public rawName?: string) {
    this.#name = rawName ?? '(no name)';
  }

  log() {
    console.log(this.#name);
  }
}

// try catch

function somethingRisky() {}

try {
  somethingRisky();
} catch (err: unknown) {
  if (err instanceof Error) {
    console.log(err.stack);
  } else {
    console.log(err);
  }
}

// // assertion in typescript

function assertIsError(err: any): asserts err is Error {
  if (!(err instanceof Error)) throw new Error(`Not an error ${err}`);
}

try {
  somethingRisky();
} catch (err: unknown) {
  assertIsError(err);
  console.log(err.stack);
}
