package test.fixtures.code_samples;

class HelloWorld
{
    public static void main(String args[])
    {
        /* This is a block comment
        stating that we are performing a simple operation
        of printing hello world below*/
        
        String s = "Hello, World!";
        System.out.println(s);
        
        // simple arithmetic operations
        int a = 5;
        int b = 7;

        System.out.printf("sum of two numbers is: %d \n", a + b);
        System.out.printf("difference of two numbers is: %d \n", a - b);
        System.out.printf("product of two numbers is: %d \n", a * b);
        System.out.printf("quotient of two numbers is: %d \n", a / b);

    }
}