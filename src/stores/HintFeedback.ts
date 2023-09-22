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
