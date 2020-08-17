DROP TABLE IF EXISTS quote;
CREATE TABLE quote
(
    id    int(12) unsigned NOT NULL AUTO_INCREMENT,
    quote varchar(255) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `quote`
--

LOCK TABLES 'quote' WRITE;
/*!40000 ALTER TABLE `quote` DISABLE KEYS */;
INSERT INTO quote
VALUES
    (1, 'Working with Spring Boot is like pair-programming with the Spring developers.'),
    (2, 'With Boot you deploy everywhere you can find a JVM basically.'),
    (3, 'Spring has come quite a ways in addressing developer enjoyment and ease of use since the last time I built an application using it.'),
    (4, 'Previous to Spring Boot, I remember XML hell, confusing set up, and many hours of frustration.'),
    (5, 'Spring Boot solves this problem. It gets rid of XML and wires up common components for me, so I don\' t have to spend hours scratching my head just to figure out how it\'s all pieced together.'),
    (6, 'It embraces convention over configuration, providing an experience on par with frameworks that excel at early stage development, such as Ruby on Rails.'),
    (7, 'The real benefit of Boot, however, is that it\' s just Spring.That means any direction the code takes, regardless of complexity, I know it\'s a safe bet.'),
    (8, 'I don\' t worry about my code scaling.Boot allows the developer to peel back the layers and customize when it\'s appropriate while keeping the conventions that just work.'),
    (9, 'So easy it is to switch container in #springboot.'),
    (10, 'Really loving Spring Boot, makes stand alone Spring apps easy.'),
    (11, 'I have two hours today to build an app from scratch. @springboot to the rescue!'),
    (12, '@springboot with @springframework is pure productivity! Who said in #java one has to write double the code than in other langs? #newFavLib');
/*!40000 ALTER TABLE `quote` ENABLE KEYS */;
UNLOCK TABLES;
