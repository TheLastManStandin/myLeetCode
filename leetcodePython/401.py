class Watch:
    def __init__(self, leds: list):
        h_leds = leds[6:]
        m_leds = leds[:6]
        self.hour = 0
        self.minute = 0

        for h_led_index in range(len(h_leds)):
            self.hour += h_leds[h_led_index] * (2 ** h_led_index)
        
        self.hour = self.hour

        for m_led_index in range(len(m_leds)):
            self.minute += m_leds[m_led_index] * (2 ** m_led_index)

        self.minute = self.minute

    def get_normal_time(self):
        if self.hour < 12 and self.minute < 60:
            return [self.hour, self.minute]
        else: 
            return None

class Solution:
    def __init__(self):
        self.out = list()

    def recursion(self, turnedOn, position, to_turn_on, massive):
        if to_turn_on > 0:
            for i in range(position, 10):
                massive[i] = 1
                self.recursion(turnedOn, i + 1, to_turn_on - 1, massive)
                massive[i] = 0
        else:
            w = Watch(massive)
            # print(massive)
            # print(w.get_normal_time())
            w_time = w.get_normal_time() 
            if w_time:
                self.out.append(w_time)
            return

    def readBinaryWatch(self, turnedOn: int) -> list[str]:
        massive = [0 for _ in range(10)]
        out = list()                    

        self.recursion(turnedOn, 0, turnedOn, massive)

        # w = Watch([0, 0, 1, 0, 1, 1, 0, 0, 1, 1])
        # print(w.get_normal_time())
        formated_out = list()

        for time in sorted(out):
            if time[1] < 10:
                time[1] = "0" + str(time[1])
            else:
                self.minute = str(time[1])

            
            formated_out.append(str(time[0]) + ":" + str(time[1]))

        return formated_out

s = Solution()
print(s.readBinaryWatch(2))