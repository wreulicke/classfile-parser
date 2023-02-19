package main;

import java.lang.annotation.ElementType;
import java.lang.annotation.Repeatable;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Deprecated
class Test{

    public static final int CONSTANT = 0;
    private final String x = "test";

    @Deprecated
    @Annot
    public static void main(@Annot String[] args) throws @Annot Exception {
        @Annot long x = 25;
        System.out.println(x);
    }

    @Annot
    static class StaticNested {}
    
    @Annot
    class Nested {}

    record Point(@Annot int x, int y) { }
}

class TypeParameter<@Annot T> extends @Annot Test {

    @Annot("test")
    public <@Annot t> @Annot T typeParameter() {
        return null;
    }
}

@Target({
    ElementType.TYPE_USE,
    ElementType.TYPE,
    ElementType.TYPE_PARAMETER,
    ElementType.LOCAL_VARIABLE,
    ElementType.CONSTRUCTOR,
    ElementType.ANNOTATION_TYPE,
    ElementType.PARAMETER,
    ElementType.PACKAGE,
    ElementType.MODULE,
    ElementType.METHOD,
})
@Repeatable(Annots.class)
@Retention(RetentionPolicy.RUNTIME)
@interface Annot {
    String value() default "default";
}

@Target({
    ElementType.TYPE_USE,
    ElementType.TYPE,
    ElementType.TYPE_PARAMETER,
    ElementType.LOCAL_VARIABLE,
    ElementType.CONSTRUCTOR,
    ElementType.ANNOTATION_TYPE,
    ElementType.PARAMETER,
    ElementType.PACKAGE,
    ElementType.MODULE,
    ElementType.METHOD,
})
@Retention(RetentionPolicy.RUNTIME)
@interface Annots {
    Annot[] value();
}