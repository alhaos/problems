from manim import *

class SwapCircular(Scene):
    def construct(self):
        values = [5, 2, 8, 1, 6]

        boxes = [Square(side_length=1.5).move_to(LEFT * 3 + RIGHT * i * 2) for i in range(len(values))]
        texts = [Text(str(values[i])).scale(0.7).move_to(boxes[i].get_center()) for i in range(len(values))]

        for box, text in zip(boxes, texts):
            self.add(box, text)

        self.wait(1)

        i, j = 1, 3

        self.play(
            boxes[i].animate.shift(UP * 3.5).shift(RIGHT * (j - i) * 2).shift(DOWN * 3.5),
            boxes[j].animate.shift(DOWN * 3.5).shift(LEFT * (j - i) * 2).shift(UP * 3.5),
            texts[i].animate.shift(UP * 3.5).shift(RIGHT * (j - i) * 2).shift(DOWN * 3.5),
            texts[j].animate.shift(DOWN * 3.5).shift(LEFT * (j - i) * 2).shift(UP * 3.5),
            run_time=1.5
        )

        boxes[i], boxes[j] = boxes[j], boxes[i]
        texts[i], texts[j] = texts[j], texts[i]

        self.wait(2)
