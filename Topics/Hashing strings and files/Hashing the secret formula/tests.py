import os

from hstest import StageTest, CheckResult, WrongAnswer, TestCase

SHA512_output = ["090ed565b68782cf59752744201d6fb860423854e568eec76cff2f87bc817f817370d7a977a3e379e0f493590c87c543e5401d68ba61f363a582ba043a375889"]

FILENAME = "secret_formula.txt"


class TestAdmissionProcedure(StageTest):
    def generate(self):
        return [TestCase(stdin=[test], attach=[test]) for test in SHA512_output]

    def check(self, reply: str, clue: str):
        if not os.path.exists(FILENAME):
            raise WrongAnswer(f"Cannot find file {FILENAME}")

        if clue[0].rstrip() != reply.rstrip():
            raise WrongAnswer(
                f"Incorrect! 😵❌ Wrong answer!\n"
                f"Your program printed:\n{reply.rstrip()}\n"
                f"\nAre you sure you properly opened {FILENAME}?")

        print(f"\nWell done! Mr. Krabs says thanks!\n")
        return CheckResult.correct()

if __name__ == '__main__':
    TestAdmissionProcedure().run_tests()