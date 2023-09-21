import { onDestroy } from 'svelte';
import { writable, type Unsubscriber, type Writable, get } from 'svelte/store';
// import { derived, type Readable } from 'svelte/store';

export const hoveredHint = writable(false);

const INIT_VISIBILITY = -62;

export const hintVisibility = writable(INIT_VISIBILITY);

export const resetVisibility = () => {
  return hintVisibility.set(INIT_VISIBILITY);
}

export const hintShown = () => {
  hintVisibility.update((v) => {
    if (v - 5 >= -72) {
      return v - 5;
    }
    return -72;
  });
}

export const hintFeedbackMessage = writable("");

interface timeI {
  time: number
};
interface mutexI {
  mutex: boolean
}
interface messageI {
  message: string
}

class HintFeedback {
  queue: string[];
  queueWritable: Writable<string[]>;
  mutexWritable: Writable<mutexI>;
  waitTimeWritable: Writable<timeI>;
  message: Writable<messageI>;
  unsubscribers: Unsubscriber[];

  constructor() {
    this.queue = [];
    this.queueWritable = writable<string[]>([]);
    this.mutexWritable = writable({ mutex: false});
    this.waitTimeWritable = writable<timeI>();
    this.message = writable<messageI>();

    const unsubscribe = this.queueWritable.subscribe((q: string[]) => {
      if (q.length > 0) {

        if (!this.isLocked() && q.length > 0) {
          this.lock();

          const message = this.pop();
          this.message.set({ message: message });
          console.log("unsubscribe: set the message:", message);
        } else {
          console.log("unsubscribe: didn't match the !this.isLocked() && q.length > 0;", `isLocked:${this.isLocked()} , q.length:${q.length}`);
        }
      } else {
        console.log("unsubscribe: empty update:", q);
      }
    });

    const unsubscribe1 = this.mutexWritable.subscribe(() => {
      this.queueWritable.update(q => q);
      console.log("unsubscribe1: updated the mutex");
    });

    const unsubscribe2 = this.waitTimeWritable.subscribe((time: timeI) => {
      if (time) {
        this.unlockIn(time);
        console.log("unsubscribe2: updated for time:", time);
        this.queueWritable.update(q => q);
      }
    });

    this.unsubscribers = [unsubscribe, unsubscribe1, unsubscribe2];
  }

  lock() {
    this.mutexWritable.set({ mutex: true });
  }

  unlockIn(n: timeI) {
    const { time } = n;
    const t = setTimeout(() => {
      console.log("unlocked after:", n);
      this.mutexWritable.set({ mutex: false });
      clearTimeout(t);
    }, time);
  }

  isLocked(): boolean {
    console.log("isLocked(): ", get(this.mutexWritable));
    return get(this.mutexWritable).mutex;
  }

  push(m: string) {
    this.queueWritable.update(q => {
      q.push(m);

      console.log("pushed to message:", m, "q:", q);
      
      return q;
    });
  }

  pop(): string {
    let out: string = "";
    this.queueWritable.update(q => {

      if (q.length > 0)  {
        const t = q.splice(0, 1);
        out = t[0];
      }

      return q;
    });

    console.log(`poped(): ${out}`);

    return out;
  }

  onDestroy() {
    this.unsubscribers.forEach((unsub: Unsubscriber) => onDestroy(unsub));
  }
}

export const hintFeedback = new HintFeedback();

// const test1 = ["qwer qwer qwer", "asdf asdf asdf", "zxcv zxcv zxcv"].reverse();
// test1.forEach((v) => {
//   hintFeedback.push(v);
// })


// const hintFeedbackMessageQuque = writable<string[]>([]);

// export const feedbackMessagePush = (m: string) => {
//   hintFeedbackMessageQuque.update(s => {
//     s.push(m);
//     return s;
//   })
// }

// export const feedbackMessagePop = () => {
//   let out: string = "";
//   hintFeedbackMessageQuque.update(s => {
//     if (s.length != 0)  {
//       const t = s.slice(0, 1);
//       if (t.length == 1) {
//         out = t[0];
//       }
//     }
//     return s;
//   })
//   return out;
// }

// const testNotificationIsTheSameArrayWerePut = () => {
//   // const store = writable<{a: number}>({ a: 0});
//   const store = writable<string[]>([]);
//   // const store = writable<boolean>(true);
//   // store.subscribe((messages: boolean) => {
//   store.subscribe((messages: string[]) => {
//   // store.subscribe((messages: {a: number}) => {
//   // store.subscribe((messages: string) => {
//     console.log("test:", messages);
//   })

//   console.log("test: get:", get(store));


//   // const val = get(store);
//   // console.log("get:", val);
//   store.set(["a"]);
//   store.set(["a"]);
//   store.update(() => ["a"]);
//   store.update(() => ["a"]);
//   // store.set("a");
//   // store.set("a");
//   // store.update(() => "m");
//   // store.update(() => "m");
//   // store.update(() => true);
//   // store.update(() => true);

//   // store.set([1]);
//   // console.log("test: get:", get(store));
//   // store.set([1]);

//   // store.set({ a: 1});
//   // console.log("test: get:", get(store));
//   // store.set({ a: 1});
// }
// testNotificationIsTheSameArrayWerePut()
