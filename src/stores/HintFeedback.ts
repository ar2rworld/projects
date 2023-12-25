import { onDestroy } from 'svelte';
import { writable, type Unsubscriber, type Writable, get } from 'svelte/store';

export const hoveredHint = writable(false);

const INIT_VISIBILITY = -62;

export const hintVisibility = writable(INIT_VISIBILITY);

export const resetVisibility = () => {
	return hintVisibility.set(INIT_VISIBILITY);
};

export const hintShown = () => {
	hintVisibility.update((v) => {
		if (v - 5 >= -72) {
			return v - 5;
		}
		return -72;
	});
};

export const hintFeedbackMessage = writable('');

interface timeI {
	time: number;
}
interface mutexI {
	mutex: boolean;
}
interface messageI {
	message: string;
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
		this.mutexWritable = writable({ mutex: false });
		this.waitTimeWritable = writable<timeI>();
		this.message = writable<messageI>();

		const unsubscribe = this.queueWritable.subscribe((q: string[]) => {
			if (q.length == 0) {
				return;
			}

			if (!this.isLocked()) {
				this.lock();

				const message = this.pop();
				this.message.set({ message: message });
			}
		});

		const unsubscribe1 = this.mutexWritable.subscribe(() => {
			this.queueWritable.update((q) => q);
		});

		const unsubscribe2 = this.waitTimeWritable.subscribe((time: timeI) => {
			if (time) {
				this.unlockIn(time);
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
			this.mutexWritable.set({ mutex: false });

			clearTimeout(t);
		}, time);
	}

	isLocked(): boolean {
		return get(this.mutexWritable).mutex;
	}

	push(m: string) {
		this.queueWritable.update((q) => {
			q.push(m);

			return q;
		});
	}

	pop(): string {
		let out: string = '';
		this.queueWritable.update((q) => {
			if (q.length > 0) {
				const t = q.splice(0, 1);
				out = t[0];
			}

			return q;
		});

		return out;
	}

	onDestroy() {
		this.unsubscribers.forEach((unsub: Unsubscriber) => onDestroy(unsub));
	}
}

export const hintFeedback = new HintFeedback();
